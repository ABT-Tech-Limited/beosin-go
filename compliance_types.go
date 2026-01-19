package beosin

// DepositRequest represents a request for deposit transaction assessment
type DepositRequest struct {
	// ChainID is the blockchain chain ID
	ChainID string `json:"chainId"`

	// Hash is the transaction hash
	Hash string `json:"hash"`

	// Token is the token address (optional, for native tokens use token name)
	Token string `json:"token,omitempty"`
}

// WithdrawalRequest represents a request for withdrawal transaction assessment
type WithdrawalRequest struct {
	// ChainID is the blockchain chain ID
	ChainID string `json:"chainId"`

	// Hash is the transaction hash
	Hash string `json:"hash"`

	// Token is the token address (optional, for native tokens use token name)
	Token string `json:"token,omitempty"`
}

// AddressRiskRequest represents a request for address risk assessment
type AddressRiskRequest struct {
	// ChainID is the blockchain chain ID
	ChainID string `json:"chainId"`

	// Address is the address hash
	Address string `json:"address"`

	// Token is the token address (optional, defaults to native token)
	Token string `json:"token,omitempty"`
}

// MaliciousAddressRequest represents a request for malicious address query
type MaliciousAddressRequest struct {
	// ChainID is the blockchain chain ID
	ChainID string `json:"chainId"`

	// Address is the address hash
	Address string `json:"address"`
}

// VASPRequest represents a request for VASP query
type VASPRequest struct {
	// ChainID is the blockchain chain ID
	ChainID string `json:"chainId"`

	// Address is the address hash
	Address string `json:"address"`
}

// RiskDetail represents a single risk detail
type RiskDetail struct {
	// RiskName is the name of the risk type
	RiskName string `json:"riskName"`

	// Rate is the proportion of funds (4 decimal places)
	Rate float64 `json:"rate"`

	// Amount is the amount involved
	Amount float64 `json:"amount"`
}

// Risk represents a risk item with strategy and details
type Risk struct {
	// RiskStrategy is the risk strategy type
	RiskStrategy string `json:"riskStrategy"`

	// RiskDetails contains the details of the risk
	RiskDetails []RiskDetail `json:"riskDetails"`
}

// TransactionRiskData represents the data in transaction risk response
type TransactionRiskData struct {
	// Score is the risk score
	Score float64 `json:"score"`

	// RiskLevel is the risk level (Severe/High/Medium/Low)
	RiskLevel string `json:"riskLevel"`

	// Risks contains the list of detected risks
	Risks []Risk `json:"risks"`
}

// TransactionRiskResponse represents the response from transaction assessment
type TransactionRiskResponse struct {
	BaseResponse
	Data *TransactionRiskData `json:"data"`
}

// StrategyRiskDetail represents risk details for a strategy
type StrategyRiskDetail struct {
	// StrategyName is the name of the risk strategy
	StrategyName string `json:"strategyName"`

	// RiskDetails contains the details of the risk
	RiskDetails []RiskDetail `json:"riskDetails"`
}

// AddressRiskData represents the data in address risk response
type AddressRiskData struct {
	// Score is the overall risk score (0-100)
	Score float64 `json:"score"`

	// RiskLevel is the overall risk level
	RiskLevel string `json:"riskLevel"`

	// IncomingScore is the deposit risk score
	IncomingScore float64 `json:"incomingScore"`

	// IncomingLevel is the deposit risk level
	IncomingLevel string `json:"incomingLevel"`

	// IncomingDetail contains deposit risk details
	IncomingDetail []StrategyRiskDetail `json:"incomingDetail"`

	// OutgoingScore is the withdrawal risk score
	OutgoingScore float64 `json:"outgoingScore"`

	// OutgoingLevel is the withdrawal risk level
	OutgoingLevel string `json:"outgoingLevel"`

	// OutgoingDetail contains withdrawal risk details
	OutgoingDetail []StrategyRiskDetail `json:"outgoingDetail"`

	// RiskTagScore is the risk tag score
	RiskTagScore float64 `json:"riskTagScore"`

	// RiskTagLevel is the risk tag level
	RiskTagLevel string `json:"riskTagLevel"`

	// RiskTagDetails contains risk tag types
	RiskTagDetails []string `json:"riskTagDetails"`
}

// AddressRiskResponse represents the response from address risk assessment
type AddressRiskResponse struct {
	BaseResponse
	Data *AddressRiskData `json:"data"`
}

// MaliceTag represents a malice tag
type MaliceTag struct {
	// TagType is the tag type
	TagType string `json:"tagType"`

	// Tag is the tag value
	Tag string `json:"tag"`
}

// MaliceDetail represents malice details
type MaliceDetail struct {
	// Source is the source of the malicious address
	Source string `json:"source"`

	// MaliceTags contains the malice tags
	MaliceTags []MaliceTag `json:"maliceTags"`
}

// SanctionDetail represents sanction details
type SanctionDetail struct {
	// Standard is the sanction standard (e.g., OFAC)
	Standard string `json:"standard"`

	// Tag is the sanction tag
	Tag string `json:"tag"`

	// Entity is the sanctioned entity
	Entity string `json:"entity"`

	// Country is the sanctioned country
	Country string `json:"country"`

	// Source is the source URL
	Source string `json:"source"`
}

// MaliciousAddressData represents the data in malicious address response
type MaliciousAddressData struct {
	// Address is the queried address
	Address string `json:"address"`

	// IsMalicious indicates if the address is malicious
	IsMalicious bool `json:"isMalicious"`

	// MaliceDetail contains malice details
	MaliceDetail *MaliceDetail `json:"maliceDetail"`

	// IsSanction indicates if the address is sanctioned
	IsSanction bool `json:"isSanction"`

	// SanctionDetail contains sanction details
	SanctionDetail *SanctionDetail `json:"sanctionDetail"`

	// IsInCustomerBlackList indicates if the address is in customer blacklist
	IsInCustomerBlackList bool `json:"isInCustomerBlackList"`
}

// MaliciousAddressResponse represents the response from malicious address query
type MaliciousAddressResponse struct {
	BaseResponse
	Data *MaliciousAddressData `json:"data"`
}

// VASPData represents the data in VASP response
type VASPData struct {
	// Address is the queried address
	Address string `json:"address"`

	// IsVasp indicates if the address is a VASP entity
	IsVasp bool `json:"isVasp"`

	// VaspTags contains the VASP entity tags
	VaspTags []string `json:"vaspTags"`
}

// VASPResponse represents the response from VASP query
type VASPResponse struct {
	BaseResponse
	Data *VASPData `json:"data"`
}
