package beosin

// AccountBalanceData represents the data in account balance response
type AccountBalanceData struct {
	// SurplusIntegral is the remaining credits
	SurplusIntegral int64 `json:"surplusIntegral"`

	// EquityStartDate is the equity start date as Unix timestamp
	EquityStartDate int64 `json:"equityStartDate"`

	// EquityEndDate is the equity end date as Unix timestamp
	EquityEndDate int64 `json:"equityEndDate"`
}

// AccountBalanceResponse represents the response from account balance query
type AccountBalanceResponse struct {
	BaseResponse
	Data *AccountBalanceData `json:"data"`
}
