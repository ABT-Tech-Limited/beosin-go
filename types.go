package beosin

// Chain ID constants for commonly used blockchains
const (
	// Full Query supported chains
	ChainBTC       = "0"
	ChainETH       = "1"
	ChainOptimism  = "10"
	ChainBSC       = "56"
	ChainTron      = "79"
	ChainPolygon   = "137"
	ChainHsk       = "177"
	ChainLTC       = "227"
	ChainZksync    = "324"
	ChainIoTeX     = "4689"
	ChainKaia      = "8217"
	ChainArbitrum  = "42161"
	ChainAvalanche = "43114"
	ChainAptos     = "aptos"
	ChainSolana    = "solana"
	ChainTON       = "ton"
	ChainXRP       = "xrp"

	// Basic Query supported chains
	ChainBase         = "8453"
	ChainLinea        = "59144"
	ChainScroll       = "534352"
	ChainMerlin       = "4200"
	ChainNeo          = "888"
	ChainZklink       = "810180"
	ChainRonin        = "2020"
	ChainBerachain    = "80084"
	ChainMonad        = "monad"
	ChainAstar        = "592"
	ChainTaiko        = "167000"
	ChainBitlayer     = "200901"
	ChainSui          = "sui"
	ChainSei          = "1329"
	ChainKCC          = "321"
	ChainSonic        = "146"
	ChainConfluxESpace = "1030"
)

// Risk levels returned by the API
const (
	RiskLevelSevere = "Severe"
	RiskLevelHigh   = "High"
	RiskLevelMedium = "Medium"
	RiskLevelLow    = "Low"
)

// Exposure types for V4 API
const (
	ExposureDirect   = "Direct"
	ExposureIndirect = "Indirect"
)

// BaseResponse represents the common response structure from Beosin API
type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// IsSuccess checks if the API response indicates success
func (r *BaseResponse) IsSuccess() bool {
	return r.Code == 200
}
