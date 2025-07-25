package collector

import (
	"net/http"
)

const (
	// kubeletSummaryEndpoint is the endpoint for the kubelet summary API.
	kubeletSummaryEndpoint = "https://localhost:10250/stats/summary"
)

// buildHTTPRequest creates a new HTTP request to the kubelet summary endpoint
// with the necessary authorization header.
func buildHTTPRequest(endpoint string) (*http.Request, error) {
	if endpoint == "" {
		endpoint = kubeletSummaryEndpoint
	}

	// create a new HTTP request to the kubelet summary endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}
