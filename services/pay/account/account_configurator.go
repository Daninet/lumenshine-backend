package account

import (
	"fmt"
	"github.com/Soneso/lumenshine-backend/helpers"
	cerr "github.com/Soneso/lumenshine-backend/icop_error"
	"github.com/Soneso/lumenshine-backend/services/pay/config"
	"github.com/Soneso/lumenshine-backend/services/pay/db"
	"github.com/Soneso/lumenshine-backend/services/pay/stellar"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	m "github.com/Soneso/lumenshine-backend/services/db/models"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/support/log"
	"github.com/stellar/go/xdr"
)

// Configurator is responsible for configuring new Stellar accounts that
// participate in ICO.
type Configurator struct {
	DB      *db.DB
	log     *logrus.Entry
	Client  *rpcclient.Client
	cnf     *config.Config
	Horizon *horizon.Client
}

//NewAccountConfigurator initialises the account configurator
func NewAccountConfigurator(DB *db.DB, cnf *config.Config) *Configurator {
	var err error

	l := new(Configurator)
	l.DB = DB
	l.cnf = cnf
	l.log = helpers.GetDefaultLog("Account-Configurator", "")

	_, err = keypair.Parse(cnf.Stellar.IssuerPublicKey)
	if err != nil || (err == nil && cnf.Stellar.IssuerPublicKey[0] != 'G') {
		log.WithField("err", err).Error("Invalid IssuerPublicKey")
		os.Exit(-1)
	}

	_, err = keypair.Parse(cnf.Stellar.DistributionPublicKey)
	if err != nil || (err == nil && cnf.Stellar.DistributionPublicKey[0] != 'G') {
		log.WithField("err", err).Error("Invalid DistributionPublicKey")
		os.Exit(-1)
	}

	l.Horizon = &horizon.Client{
		URL: cnf.Stellar.Horizon,
		HTTP: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	root, err := l.Horizon.Root()
	if err != nil {
		log.WithField("err", err).Error("Error loading Horizon root")
		os.Exit(-1)
	}

	if root.NetworkPassphrase != cnf.Stellar.NetworkPassphrase {
		log.Errorf("Invalid network passphrase (have=%s, want=%s)", root.NetworkPassphrase, cnf.Stellar.NetworkPassphrase)
		os.Exit(-1)
	}

	l.log.Info("Accountconfigurator created")
	return l
}

func (c *Configurator) hasTrustline(acc *horizon.Account) bool {
	//check if trustline present
	for _, b := range acc.Balances {
		if b.Asset.Code == c.cnf.Stellar.TokenAssetCode && b.Asset.Issuer == c.cnf.Stellar.IssuerPublicKey {
			return true
		}
	}
	return false
}

//GetTrustStatus checks if the user exists and creates him otherwise.
//returns if the trustline to the coin is present on the useraccount
func (c *Configurator) GetTrustStatus(o *m.UserOrder) (bool, error) {
	var acc horizon.Account
	var err error
	var exists bool

	acc, exists, err = c.getAccount(o.UserStellarPublicKey)
	if err != nil {
		return false, err
	}

	if !exists {
		err := c.createAccount(o.UserStellarPublicKey)
		if err != nil {
			return false, errors.Wrap(err, "Error creating account")
		}

		//need to relaod the account to get the sequence
		acc, exists, err = c.getAccount(o.UserStellarPublicKey)
		if err != nil {
			return false, err
		}
	}

	return c.hasTrustline(&acc), nil
}

//GetPaymentTransaction creates the transaction for a valid payment
//returns the transaction or an error code
func (c *Configurator) GetPaymentTransaction(o *m.UserOrder) (string, int64, error) {
	var acc horizon.Account
	var err error
	var exists bool

	acc, exists, err = c.getAccount(o.UserStellarPublicKey)
	if err != nil {
		return "", 0, errors.Wrap(err, "Could not read stellar account")
	}

	if !exists {
		return "", cerr.StellarAccountNotExists, nil
	}

	if !c.hasTrustline(&acc) {
		return "", cerr.StellarTrustlineNotExists, nil
	}

	//create the transaction and save it to DB
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: o.UserStellarPublicKey},
		build.Network{Passphrase: c.cnf.Stellar.NetworkPassphrase},
		build.AutoSequence{SequenceProvider: c.Horizon},
		build.Payment(
			build.SourceAccount{AddressOrSeed: c.cnf.Stellar.DistributionPublicKey},
			build.Destination{AddressOrSeed: o.UserStellarPublicKey},
			build.CreditAmount{
				Code:   c.cnf.Stellar.TokenAssetCode,
				Issuer: c.cnf.Stellar.IssuerPublicKey,
				Amount: fmt.Sprintf("%d", o.CoinAmount),
			},
		),
	)
	if err != nil {
		return "", 0, errors.Wrap(err, "Could not create transaction")
	}

	var txe build.TransactionEnvelopeBuilder
	err = txe.Mutate(tx)
	if err != nil {
		return "", 0, errors.Wrap(err, "Could not create TransactionEnvelopeBuilder")
	}

	txes, err := txe.Base64()
	if err != nil {
		return "", 0, errors.Wrap(err, "Could not bas64 encode TransactionEnvelopeBuilder")
	}

	return txes, 0, nil
}

//ExecuteTransaction checks the transaction, signs it and executes it
func (c *Configurator) ExecuteTransaction(o *m.UserOrder, tx string) error {
	/*var acc horizon.Account
	var err error
	var exists bool*/

	txe, err := c.decodeFromBase64(tx)
	if err != nil {
		return errors.Wrap(err, "Could not decode transaction")
	}
	//check the operations inside the transaction
	if txe.E.Tx.SourceAccount.Address() != o.UserStellarPublicKey {
		return errors.New("Source not like order")
	}
	if len(txe.E.Tx.Operations) != 1 {
		return errors.New("Operation-Count not 1")
	}
	op := txe.E.Tx.Operations[0].Body.PaymentOp
	if op == nil {
		return errors.New("No Payment operation in tx")
	}
	if op.Destination.Address() != o.UserStellarPublicKey {
		return errors.New("Destination not like order")
	}
	amount, err := stellar.StroopToCoin(int64(op.Amount))
	if err != nil {
		return errors.Wrap(err, "Amount to small")
	}
	if amount.Int64() != o.CoinAmount {
		return errors.New("Amount does not match")
	}

	opIssuer := ""
	opAssetCode := ""
	if op.Asset.Type == xdr.AssetTypeAssetTypeCreditAlphanum12 {
		opIssuer = op.Asset.AlphaNum12.Issuer.Address()
		opAssetCode = string(op.Asset.AlphaNum12.AssetCode[:len(c.cnf.Stellar.TokenAssetCode)])
	} else if op.Asset.Type == xdr.AssetTypeAssetTypeCreditAlphanum4 {
		opIssuer = op.Asset.AlphaNum4.Issuer.Address()
		opAssetCode = string(op.Asset.AlphaNum4.AssetCode[:len(c.cnf.Stellar.TokenAssetCode)])
	} else {
		return errors.New("Wrong AssetType")
	}

	if opAssetCode != c.cnf.Stellar.TokenAssetCode {
		return errors.New("Wrong asset code")
	}

	if opIssuer != c.cnf.Stellar.IssuerPublicKey {
		return errors.New("Wrong issuer")
	}

	//check signer exact 2, check public keys of user and dist in signer only

	/*txes, err := txe.Base64()
	if err != nil {
		return errors.Wrap(err, "Could not base64 decode transaction")
	}
	if txes != o.PaymentTX {
		return errors.Wrap(err, "Transaction does not match saved transaction")
	}*/

	//sign the transaxtion with the dest seed
	s := build.Sign{Seed: c.cnf.Stellar.DistributionSeed}
	err = s.MutateTransactionEnvelope(txe)
	if err != nil {
		return errors.Wrap(err, "Could not sign")
	}

	txeBase64, err := txe.Base64()
	if err != nil {
		return errors.Wrap(err, "Could not base64 decode txe")
	}

	return c.submitTransaction(txeBase64)
}

//createAccount create the user stellar account
//it uses the configured distribution account as source and uses autosequence
func (c *Configurator) createAccount(account string) error {
	tx, err := c.buildTransaction(
		c.cnf.Stellar.DistributionPublicKey,
		c.cnf.Stellar.DistributionSeed,
		0,
		build.CreateAccount(
			build.SourceAccount{AddressOrSeed: c.cnf.Stellar.DistributionPublicKey},
			build.Destination{AddressOrSeed: account},
			build.NativeAmount{Amount: c.cnf.Stellar.StartingBalance},
		),
	)

	if err != nil {
		return errors.Wrap(err, "Error building user create transaction")
	}

	err = c.submitTransaction(tx)
	if err != nil {
		return errors.Wrap(err, "Error creating user stellar account")
	}

	return nil
}

func (c *Configurator) getAccount(account string) (horizon.Account, bool, error) {
	var hAccount horizon.Account
	hAccount, err := c.Horizon.LoadAccount(account)
	if err != nil {
		if err, ok := err.(*horizon.Error); ok && err.Response.StatusCode == http.StatusNotFound {
			return hAccount, false, nil
		}
		return hAccount, false, err
	}

	return hAccount, true, nil
}

//SignAndRunTransaction signs the transaction and runs it
//used while developing
func (c *Configurator) SignAndRunTransaction(tx string, seed string) error {
	return c.signAndSubmitTransaction(tx, seed)
}
