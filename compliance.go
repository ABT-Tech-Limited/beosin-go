package beosin

import "context"

const (
	endpointDeposit          = "/api/v2/kyt/tx/deposit"
	endpointWithdraw         = "/api/v2/kyt/tx/withdraw"
	endpointAddressRisk      = "/api/v3/kyt/address/risk"
	endpointMaliciousAddress = "/api/v2/kyt/tag/malicious"
	endpointVASP             = "/api/v2/kyt/tag/vasp"
)

// DepositTransactionAssessment performs risk assessment on deposit transactions
func (c *client) DepositTransactionAssessment(ctx context.Context, req *DepositRequest) (*TransactionRiskResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"hash":    req.Hash,
		"token":   req.Token,
	})

	var resp TransactionRiskResponse
	if err := c.doRequest(ctx, endpointDeposit, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// WithdrawalTransactionAssessment performs risk assessment on withdrawal transactions
func (c *client) WithdrawalTransactionAssessment(ctx context.Context, req *WithdrawalRequest) (*TransactionRiskResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"hash":    req.Hash,
		"token":   req.Token,
	})

	var resp TransactionRiskResponse
	if err := c.doRequest(ctx, endpointWithdraw, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EOAAddressRiskAssessment performs risk assessment on EOA addresses
func (c *client) EOAAddressRiskAssessment(ctx context.Context, req *AddressRiskRequest) (*AddressRiskResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"address": req.Address,
		"token":   req.Token,
	})

	var resp AddressRiskResponse
	if err := c.doRequest(ctx, endpointAddressRisk, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MaliciousAddressQuery queries if an address is malicious
func (c *client) MaliciousAddressQuery(ctx context.Context, req *MaliciousAddressRequest) (*MaliciousAddressResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"address": req.Address,
	})

	var resp MaliciousAddressResponse
	if err := c.doRequest(ctx, endpointMaliciousAddress, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VASPQuery queries if an address is a VASP entity
func (c *client) VASPQuery(ctx context.Context, req *VASPRequest) (*VASPResponse, error) {
	params := buildQueryParams(map[string]string{
		"chainId": req.ChainID,
		"address": req.Address,
	})

	var resp VASPResponse
	if err := c.doRequest(ctx, endpointVASP, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
