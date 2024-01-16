package check

import (
	"context"
	"net/http"
	"strings"
)

type PingResponse struct {
	Up bool `json:"up"`
}

// Ping pings a specific site and determines whether it's up or down right now.
func Ping(ctx context.Context, url string) (*PingResponse, error) {
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		url = "https://" + url
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &PingResponse{Up: false}, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return &PingResponse{Up: false}, nil
	}
	return &PingResponse{Up: true}, nil
}
