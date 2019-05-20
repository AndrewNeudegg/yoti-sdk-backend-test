package blackbox

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
)

const (
	HTTPClientTimeout = 5 // HTTPClientTimeout is a five second client timeout.
)

// HTTPGet will return the body of a given endpoint. Timeout after const client timeout.
func HTTPGet(hostURL string) (string, int, error) {
	logging.Info(fmt.Sprintf("Making GET Request to: %s", hostURL))

	timeout := time.Duration(HTTPClientTimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(hostURL)
	if err != nil {
		return "", 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), resp.StatusCode, err
}

// HTTPPost will POST content to a specific end point.
func HTTPPost(hostURL string, bodyContent string) (string, int, error) {
	var requestBodyBytes = []byte(bodyContent)
	req, err := http.NewRequest("POST", hostURL, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", 0, err
	}
	timeout := time.Duration(HTTPClientTimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	return string(body), resp.StatusCode, err
}

// HandleError is a utility method for tests that will fail a given
// test if an error is detected.
func HandleError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}

// testHeader shows before a test starts, this helps to differentiate different log messages.
func testHeader(t *testing.T) {
	logging.Info(fmt.Sprintf("=============== Starting test [%s] ===============", t.Name()))
}

// testFooter shows before a test starts, this helps to differentiate different log messages.
func testFooter(t *testing.T) {
	logging.Info(fmt.Sprintf("=============== Ending test   [%s] ===============", t.Name()))
}
