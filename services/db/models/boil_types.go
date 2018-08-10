// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(wl, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	for _, w := range wl {
		buf.WriteString(w)
	}
	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

// Enum values for chain
const (
	ChainFiat = "fiat"
	ChainBTC  = "btc"
	ChainEth  = "eth"
	ChainXLM  = "xlm"
)

// Enum values for message_type
const (
	MessageTypeIos     = "ios"
	MessageTypeAndroid = "android"
	MessageTypeMail    = "mail"
)

// Enum values for mail_content_type
const (
	MailContentTypeText = "text"
	MailContentTypeHTML = "html"
)

// Enum values for notification_status_code
const (
	NotificationStatusCodeNew     = "new"
	NotificationStatusCodeSuccess = "success"
	NotificationStatusCodeError   = "error"
)

// Enum values for transaction_status
const (
	TransactionStatusNew       = "new"
	TransactionStatusProcessed = "processed"
)

// Enum values for order_status
const (
	OrderStatusWaitingForPayment = "waiting_for_payment"
	OrderStatusPaymentReceived   = "payment_received"
	OrderStatusWaitingUserTX     = "waiting_user_tx"
	OrderStatusTXCreated         = "tx_created"
	OrderStatusPaymentError      = "payment_error"
	OrderStatusFinished          = "finished"
	OrderStatusError             = "error"
	OrderStatusUnderPay          = "under_pay"
	OrderStatusOverPay           = "over_pay"
	OrderStatusNoCoinsLeft       = "no_coins_left"
	OrderStatusPhaseExpired      = "phase_expired"
)

// Enum values for payment_state
const (
	PaymentStateOpen  = "open"
	PaymentStateClose = "close"
)

// Enum values for kyc_status
const (
	KycStatusNotSupported     = "not_supported"
	KycStatusWaitingForData   = "waiting_for_data"
	KycStatusWaitingForReview = "waiting_for_review"
	KycStatusInReview         = "in_review"
	KycStatusPending          = "pending"
	KycStatusRejected         = "rejected"
	KycStatusApproved         = "approved"
)

// Enum values for device_type
const (
	DeviceTypeApple  = "apple"
	DeviceTypeGoogle = "google"
)
