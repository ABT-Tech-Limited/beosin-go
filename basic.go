package beosin

import (
	"context"
	"net/url"
)

const (
	endpointAccountBalance = "/api/v1/package/info"
)

// GetAccountBalance queries the account balance
func (c *client) GetAccountBalance(ctx context.Context) (*AccountBalanceResponse, error) {
	var resp AccountBalanceResponse
	if err := c.doRequest(ctx, endpointAccountBalance, url.Values{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
