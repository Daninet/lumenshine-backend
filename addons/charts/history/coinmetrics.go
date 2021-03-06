package history

import (
	"github.com/Soneso/lumenshine-backend/addons/charts/config"
	"io"
	"net/http"
	"os"
	"path"
)

//downloadFile will download the file from the url location to a localpath
func downloadFile(url string) (localPath string, err error) {

	// Create the file
	localPath = config.Cnf.LocalDownloadDir + path.Base(url)
	out, err := os.Create(localPath)
	if err != nil {
		return
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}

	return
}
