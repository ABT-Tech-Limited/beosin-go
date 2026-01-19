package beosin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Client defines the interface for Beosin API operations
type Client interface {
	// Basic module
	GetAccountBalance(ctx context.Context) (*AccountBalanceResponse, error)

	// Compliance module
	DepositTransactionAssessment(ctx context.Context, req *DepositRequest) (*TransactionRiskResponse, error)
	WithdrawalTransactionAssessment(ctx context.Context, req *WithdrawalRequest) (*TransactionRiskResponse, error)
	EOAAddressRiskAssessment(ctx context.Context, req *AddressRiskRequest) (*AddressRiskResponse, error)
	MaliciousAddressQuery(ctx context.Context, req *MaliciousAddressRequest) (*MaliciousAddressResponse, error)
	VASPQuery(ctx context.Context, req *VASPRequest) (*VASPResponse, error)

	// Compliance-V4 module
	V4EOAAddressRiskAssessment(ctx context.Context, req *AddressRiskRequest) (*V4AddressRiskResponse, error)
	V4DepositTransactionAssessment(ctx context.Context, req *DepositRequest) (*V4TransactionRiskResponse, error)
	V4WithdrawalTransactionAssessment(ctx context.Context, req *WithdrawalRequest) (*V4TransactionRiskResponse, error)

	// Security module
	BlackAddressScreening(ctx context.Context, req *BlackScreeningRequest) (*BlackScreeningResponse, error)
}

// client is the default implementation of the Client interface
type client struct {
	options *Options
}

// NewClient creates a new Beosin API client
func NewClient(appID, appSecret string, opts ...Option) Client {
	options := &Options{
		AppID:     appID,
		AppSecret: appSecret,
	}

	for _, opt := range opts {
		opt(options)
	}

	options.applyDefaults()

	return &client{
		options: options,
	}
}

// doRequest performs an HTTP GET request to the specified endpoint
func (c *client) doRequest(ctx context.Context, endpoint string, params url.Values, result interface{}) error {
	// Build the full URL
	fullURL := c.options.BaseURL + endpoint
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	if c.options.Debug {
		log.Printf("[BEOSIN DEBUG] Request: GET %s\n", fullURL)
	}

	// Create the request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("APPID", c.options.AppID)
	req.Header.Set("APP-SECRET", c.options.AppSecret)

	// Execute the request
	resp, err := c.options.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if c.options.Debug {
		log.Printf("[BEOSIN DEBUG] Response: %s\n", string(body))
	}

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http status: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the base response to check for API errors
	var baseResp BaseResponse
	if err := json.Unmarshal(body, &baseResp); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if !baseResp.IsSuccess() {
		return NewAPIError(baseResp.Code, baseResp.Msg)
	}

	// Parse the full response
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to parse response data: %w", err)
	}

	return nil
}

// buildQueryParams builds URL query parameters from a map
func buildQueryParams(params map[string]string) url.Values {
	values := url.Values{}
	for k, v := range params {
		if v != "" {
			values.Set(k, v)
		}
	}
	return values
}

// trimBaseURL removes trailing slash from base URL
func trimBaseURL(baseURL string) string {
	return strings.TrimRight(baseURL, "/")
}
