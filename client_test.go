package beosin

import (
	"context"
	"os"
	"testing"
	"time"
)

// getTestClient creates a client for testing using environment variables
func getTestClient(t *testing.T) Client {
	appID := os.Getenv("BEOSIN_APP_ID")
	appSecret := os.Getenv("BEOSIN_APP_SECRET")

	if appID == "" || appSecret == "" {
		t.Skip("BEOSIN_APP_ID and BEOSIN_APP_SECRET environment variables are required")
	}

	baseURL := os.Getenv("BEOSIN_BASE_URL")
	opts := []Option{
		WithDebug(os.Getenv("BEOSIN_DEBUG") == "true"),
		WithTimeout(60 * time.Second),
	}

	if baseURL != "" {
		opts = append(opts, WithBaseURL(baseURL))
	}

	return NewClient(appID, appSecret, opts...)
}

// TestGetAccountBalance tests the account balance query
func TestGetAccountBalance(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	resp, err := client.GetAccountBalance(ctx)
	if err != nil {
		t.Fatalf("GetAccountBalance failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data == nil {
		t.Error("Expected data to be non-nil")
	} else {
		t.Logf("Account balance: %d credits", resp.Data.SurplusIntegral)
		t.Logf("Equity period: %d - %d", resp.Data.EquityStartDate, resp.Data.EquityEndDate)
	}
}

// TestBlackAddressScreening tests the black address screening
func TestBlackAddressScreening(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	// Test with a known sanctioned address
	req := &BlackScreeningRequest{
		Platform: "bsc",
		Address:  "0x3cffd56b47b7b41c56258d9c7731abadc360e073",
	}

	resp, err := client.BlackAddressScreening(ctx, req)
	if err != nil {
		t.Fatalf("BlackAddressScreening failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Sanction: %v, OfficialFreeze: %v", resp.Data.Sanction, resp.Data.OfficialFreeze)
		t.Logf("Has any risk: %v", resp.Data.HasAnyRisk())
	}
}

// TestMaliciousAddressQuery tests the malicious address query
func TestMaliciousAddressQuery(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &MaliciousAddressRequest{
		ChainID: ChainETH,
		Address: "0x901bb9583b24d97e995513c6778dc6888ab6870e",
	}

	resp, err := client.MaliciousAddressQuery(ctx, req)
	if err != nil {
		t.Fatalf("MaliciousAddressQuery failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Address: %s", resp.Data.Address)
		t.Logf("IsMalicious: %v, IsSanction: %v", resp.Data.IsMalicious, resp.Data.IsSanction)
	}
}

// TestVASPQuery tests the VASP query
func TestVASPQuery(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &VASPRequest{
		ChainID: ChainETH,
		Address: "0xec6ad3cb0e62cd7c8e75d2919f12c3195d998002",
	}

	resp, err := client.VASPQuery(ctx, req)
	if err != nil {
		t.Fatalf("VASPQuery failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Address: %s, IsVasp: %v", resp.Data.Address, resp.Data.IsVasp)
		if resp.Data.IsVasp {
			t.Logf("VASP Tags: %v", resp.Data.VaspTags)
		}
	}
}

// TestDepositTransactionAssessment tests deposit transaction assessment
func TestDepositTransactionAssessment(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &DepositRequest{
		ChainID: ChainETH,
		Hash:    "0x919aeb1d0ed579dbbe15a0a695b221c746c2b45d68553da0c203747c1255f739",
	}

	resp, err := client.DepositTransactionAssessment(ctx, req)
	if err != nil {
		// Task might be executing, which is expected
		if apiErr, ok := err.(*APIError); ok && apiErr.IsTaskExecuting() {
			t.Logf("Task is executing, retry later")
			return
		}
		t.Fatalf("DepositTransactionAssessment failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Score: %.2f, RiskLevel: %s", resp.Data.Score, resp.Data.RiskLevel)
		for _, risk := range resp.Data.Risks {
			t.Logf("Risk: %s", risk.RiskStrategy)
		}
	}
}

// TestWithdrawalTransactionAssessment tests withdrawal transaction assessment
func TestWithdrawalTransactionAssessment(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &WithdrawalRequest{
		ChainID: ChainETH,
		Hash:    "0xf65b781a0fc3cadb688d8b19a2dcc66dcb2393e635b390831e802fe240f0609b",
	}

	resp, err := client.WithdrawalTransactionAssessment(ctx, req)
	if err != nil {
		if apiErr, ok := err.(*APIError); ok && apiErr.IsTaskExecuting() {
			t.Logf("Task is executing, retry later")
			return
		}
		t.Fatalf("WithdrawalTransactionAssessment failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Score: %.2f, RiskLevel: %s", resp.Data.Score, resp.Data.RiskLevel)
	}
}

// TestEOAAddressRiskAssessment tests EOA address risk assessment
func TestEOAAddressRiskAssessment(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &AddressRiskRequest{
		ChainID: ChainETH,
		Address: "0x013b646fe54562a3ff6e3469fcc8c4efc2337656",
		Token:   "0xdac17f958d2ee523a2206206994597c13d831ec7", // USDT
	}

	resp, err := client.EOAAddressRiskAssessment(ctx, req)
	if err != nil {
		if apiErr, ok := err.(*APIError); ok && apiErr.IsTaskExecuting() {
			t.Logf("Task is executing, retry later")
			return
		}
		t.Fatalf("EOAAddressRiskAssessment failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Score: %.2f, RiskLevel: %s", resp.Data.Score, resp.Data.RiskLevel)
		t.Logf("Incoming: %.2f (%s), Outgoing: %.2f (%s)",
			resp.Data.IncomingScore, resp.Data.IncomingLevel,
			resp.Data.OutgoingScore, resp.Data.OutgoingLevel)
	}
}

// TestV4EOAAddressRiskAssessment tests V4 EOA address risk assessment
func TestV4EOAAddressRiskAssessment(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &AddressRiskRequest{
		ChainID: ChainETH,
		Address: "0xc0e80a623c9593b3c3911682d2084c4e93bea4a7",
		Token:   "0xdac17f958d2ee523A2206206994597c13d831EC7", // USDT
	}

	resp, err := client.V4EOAAddressRiskAssessment(ctx, req)
	if err != nil {
		if apiErr, ok := err.(*APIError); ok && apiErr.IsTaskExecuting() {
			t.Logf("Task is executing, retry later")
			return
		}
		t.Fatalf("V4EOAAddressRiskAssessment failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Score: %.2f, RiskLevel: %s", resp.Data.Score, resp.Data.RiskLevel)
	}
}

// TestV4DepositTransactionAssessment tests V4 deposit transaction assessment
func TestV4DepositTransactionAssessment(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &DepositRequest{
		ChainID: ChainPolygon,
		Hash:    "0xa7381d43259250e004478478df0715fcb85a95d6345a3dd097a18dbeebe40173",
		Token:   "0x8f3cf7ad23cd3cadbd9735aff958023239c6a063",
	}

	resp, err := client.V4DepositTransactionAssessment(ctx, req)
	if err != nil {
		if apiErr, ok := err.(*APIError); ok && apiErr.IsTaskExecuting() {
			t.Logf("Task is executing, retry later")
			return
		}
		t.Fatalf("V4DepositTransactionAssessment failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Score: %.2f, RiskLevel: %s", resp.Data.Score, resp.Data.RiskLevel)
		for _, risk := range resp.Data.Risks {
			t.Logf("Risk: %s, Exposure: %s, Hops: %d", risk.RiskStrategy, risk.Exposure, risk.Hops)
		}
	}
}

// TestV4WithdrawalTransactionAssessment tests V4 withdrawal transaction assessment
func TestV4WithdrawalTransactionAssessment(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	req := &WithdrawalRequest{
		ChainID: ChainAvalanche,
		Hash:    "0x1e6c99bb4d1e8b0858d84b8c24a62f23fcc327701469755955352e7ba9e7bc22",
		Token:   "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
	}

	resp, err := client.V4WithdrawalTransactionAssessment(ctx, req)
	if err != nil {
		if apiErr, ok := err.(*APIError); ok && apiErr.IsTaskExecuting() {
			t.Logf("Task is executing, retry later")
			return
		}
		t.Fatalf("V4WithdrawalTransactionAssessment failed: %v", err)
	}

	if resp.Code != 200 {
		t.Errorf("Expected code 200, got %d", resp.Code)
	}

	if resp.Data != nil {
		t.Logf("Score: %.2f, RiskLevel: %s", resp.Data.Score, resp.Data.RiskLevel)
	}
}

// TestAPIError tests API error handling
func TestAPIError(t *testing.T) {
	err := NewAPIError(40021, "Platform not supported")

	if !err.IsPlatformNotSupported() {
		t.Error("Expected IsPlatformNotSupported to return true")
	}

	expected := "beosin api error: code=40021, message=Platform not supported"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

// TestBlackScreeningDataHasAnyRisk tests the HasAnyRisk method
func TestBlackScreeningDataHasAnyRisk(t *testing.T) {
	tests := []struct {
		name     string
		data     BlackScreeningData
		expected bool
	}{
		{
			name:     "No risks",
			data:     BlackScreeningData{},
			expected: false,
		},
		{
			name:     "Sanction only",
			data:     BlackScreeningData{Sanction: true},
			expected: true,
		},
		{
			name:     "Multiple risks",
			data:     BlackScreeningData{Hacker: true, Mixing: true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.HasAnyRisk(); got != tt.expected {
				t.Errorf("HasAnyRisk() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
