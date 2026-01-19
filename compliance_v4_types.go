package beosin

// V4EntityDetail represents entity details in V4 API response
type V4EntityDetail struct {
	// EntityName is the name of the entity
	EntityName string `json:"entityName"`

	// Hops is the number of hops to the entity
	Hops int `json:"hops"`

	// PurificationAmountU is the purification amount in USD
	PurificationAmountU float64 `json:"purificationAmountU"`

	// PurificationRate is the purification rate
	PurificationRate float64 `json:"purificationRate"`
}

// V4Risk represents a risk item in V4 API response
type V4Risk struct {
	// RiskStrategy is the risk strategy type
	RiskStrategy string `json:"riskStrategy"`

	// Exposure is the exposure type (Direct/Indirect)
	Exposure string `json:"exposure"`

	// RiskLevel is the risk level for this strategy
	RiskLevel string `json:"riskLevel"`

	// Hops is the shortest hop count to risk entity
	Hops int `json:"hops"`

	// Rate is the proportion of funds (4 decimal places)
	Rate float64 `json:"rate"`

	// Amount is the risk amount
	Amount float64 `json:"amount"`

	// EntityDetails contains entity details
	EntityDetails []V4EntityDetail `json:"entityDetails"`
}

// V4TransactionRiskData represents the data in V4 transaction risk response
type V4TransactionRiskData struct {
	// Score is the risk score
	Score float64 `json:"score"`

	// RiskLevel is the overall risk level (Severe/High/Medium/Low)
	RiskLevel string `json:"riskLevel"`

	// Risks contains the list of detected risks
	Risks []V4Risk `json:"risks"`
}

// V4TransactionRiskResponse represents the response from V4 transaction assessment
type V4TransactionRiskResponse struct {
	BaseResponse
	Data *V4TransactionRiskData `json:"data"`
}

// V4StrategyDetail represents strategy details in V4 address risk response
type V4StrategyDetail struct {
	// StrategyName is the name of the risk strategy
	StrategyName string `json:"strategyName"`

	// Exposure is the exposure type (Direct/Indirect)
	Exposure string `json:"exposure"`

	// RiskLevel is the risk level for this strategy
	RiskLevel string `json:"riskLevel"`

	// Hops is the hop count
	Hops int `json:"hops"`

	// Rate is the risk fund ratio
	Rate float64 `json:"rate"`

	// Amount is the risk fund amount
	Amount float64 `json:"amount"`

	// EntityDetails contains entity details
	EntityDetails []V4EntityDetail `json:"entityDetails"`
}

// V4AddressRiskData represents the data in V4 address risk response
type V4AddressRiskData struct {
	// Score is the overall risk score
	Score float64 `json:"score"`

	// RiskLevel is the overall risk level
	RiskLevel string `json:"riskLevel"`

	// IncomingScore is the incoming risk score
	IncomingScore float64 `json:"incomingScore"`

	// IncomingLevel is the incoming risk level
	IncomingLevel string `json:"incomingLevel"`

	// IncomingDetail contains incoming risk details
	IncomingDetail []V4StrategyDetail `json:"incomingDetail"`

	// OutgoingScore is the outgoing risk score
	OutgoingScore float64 `json:"outgoingScore"`

	// OutgoingLevel is the outgoing risk level
	OutgoingLevel string `json:"outgoingLevel"`

	// OutgoingDetail contains outgoing risk details
	OutgoingDetail []V4StrategyDetail `json:"outgoingDetail"`

	// RiskTagScore is the risk tag score
	RiskTagScore float64 `json:"riskTagScore"`

	// RiskTagLevel is the risk tag level
	RiskTagLevel string `json:"riskTagLevel"`

	// RiskTagDetails contains risk tag types
	RiskTagDetails []string `json:"riskTagDetails"`
}

// V4AddressRiskResponse represents the response from V4 address risk assessment
type V4AddressRiskResponse struct {
	BaseResponse
	Data *V4AddressRiskData `json:"data"`
}
