package signapsesolutions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oking02/signapse-test/internal/core/domain"
	"github.com/oking02/signapse-test/internal/core/ports"
)

var _ ports.RawNumbersFetcher = (*Client)(nil)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	numbersAPI = "/api/numbers"
)

type Client struct {
	httpClient HTTPClient
	baseURL    string
	token      string
}

func NewClient(httpClient HTTPClient, baseURL, token string) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		token:      token,
	}
}

func (c Client) FetchRawNumbers(ctx context.Context) ([]any, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+numbersAPI,
		nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create fetch raw number request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed fetch raw number request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf(
				"unauthorized response [%d] from fetch raw number request: %w",
				resp.StatusCode, domain.ErrAuthFailed)
		}

		return nil, fmt.Errorf("error response [%d] from fetch raw number request", resp.StatusCode)
	}

	var result []any
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed decode fetch raw number response: %w", err)
	}

	return result, nil
}
