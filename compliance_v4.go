package beosin

import "context"

const (
	endpointV4AddressRisk = "/api/v4/kyt/address/risk"
	endpointV4Deposit     = "/api/v4/kyt/tx/deposit"
	endpointV4Withdraw    = "/api/v4/kyt/tx/withdraw"
)

// V4EOAAddressRiskAssessment performs V4 risk assessment on EOA addresses
func (c *client) V4EOAAddressRiskAssessment(ctx context.Context, req *AddressRiskRequest) (*V4AddressRiskResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"address": req.Address,
		"token":   req.Token,
	})

	var resp V4AddressRiskResponse
	if err := c.doRequest(ctx, endpointV4AddressRisk, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// V4DepositTransactionAssessment performs V4 risk assessment on deposit transactions
func (c *client) V4DepositTransactionAssessment(ctx context.Context, req *DepositRequest) (*V4TransactionRiskResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"hash":    req.Hash,
		"token":   req.Token,
	})

	var resp V4TransactionRiskResponse
	if err := c.doRequest(ctx, endpointV4Deposit, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// V4WithdrawalTransactionAssessment performs V4 risk assessment on withdrawal transactions
func (c *client) V4WithdrawalTransactionAssessment(ctx context.Context, req *WithdrawalRequest) (*V4TransactionRiskResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"hash":    req.Hash,
		"token":   req.Token,
	})

	var resp V4TransactionRiskResponse
	if err := c.doRequest(ctx, endpointV4Withdraw, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
