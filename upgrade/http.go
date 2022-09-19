package upgrade

import (
	"net/http"
	"time"
)

const defaultTimeout = 30 * time.Second

// newHTTPClient creates a new http client with timeout.
func newHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return client
}
