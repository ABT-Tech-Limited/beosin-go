package beosin

// BlackScreeningRequest represents a request for black address screening
type BlackScreeningRequest struct {
	// Platform is the blockchain platform (e.g., bsc, eth)
	Platform string `json:"platform"`

	// Address is the address to check
	Address string `json:"address"`
}

// BlackScreeningData represents the data in black screening response
type BlackScreeningData struct {
	// Sanction indicates if the address is related to global sanctions
	Sanction bool `json:"sanction"`

	// Scam indicates if the address is involved in financial fraud, phishing, Ponzi schemes
	Scam bool `json:"scam"`

	// Gambling indicates if the address is involved in offline gambling
	Gambling bool `json:"gambling"`

	// Darknet indicates if the address is involved in darknet transactions
	Darknet bool `json:"darknet"`

	// Theft indicates if the address is involved in theft
	Theft bool `json:"theft"`

	// Mixing indicates if the address provides mixing services
	Mixing bool `json:"mixing"`

	// Hacker indicates if the address exploits smart contract vulnerabilities
	Hacker bool `json:"hacker"`

	// Ransomware indicates if the address profits from ransomware
	Ransomware bool `json:"ransomware"`

	// Trojan indicates if the address profits from trojans
	Trojan bool `json:"trojan"`

	// ChildAbuseMaterial indicates if the address is related to child abuse material
	ChildAbuseMaterial bool `json:"childAbuseMaterial"`

	// Terrorist indicates if the address is related to terrorist organizations
	Terrorist bool `json:"terrorist"`

	// Drug indicates if the address is a drug trafficking organization
	Drug bool `json:"drug"`

	// Lawsuit indicates if the address is in a global judicial network
	Lawsuit bool `json:"lawsuit"`

	// BusinessBlackList indicates if the address is a business sanction address
	BusinessBlackList bool `json:"businessBlackList"`

	// Piracy indicates if the address profits from pirated software
	Piracy bool `json:"piracy"`

	// FraudShop indicates if the address is a fraud shop
	FraudShop bool `json:"fraudShop"`

	// UndergroundBank indicates if the address is an underground bank
	UndergroundBank bool `json:"undergroundBank"`

	// MoneyMule indicates if the address is a money laundering intermediary
	MoneyMule bool `json:"moneyMule"`

	// ProtocolPiracy indicates if the address is involved in protocol piracy
	ProtocolPiracy bool `json:"protocolPiracy"`

	// IllicitActorOrganization indicates if the address is an illicit actor organization
	IllicitActorOrganization bool `json:"illicitActorOrganization"`

	// HighRiskExchange indicates if the address is a high-risk exchange
	HighRiskExchange bool `json:"highRiskExchange"`

	// HighRiskJurisdictionFATF indicates if the address is in a FATF high-risk jurisdiction
	HighRiskJurisdictionFATF bool `json:"highRiskJurisdictionFATF"`

	// GreyListFATF indicates if the address is on the FATF grey list
	GreyListFATF bool `json:"greyListFATF"`

	// OfficialFreeze indicates if the address is officially frozen
	OfficialFreeze bool `json:"officialFreeze"`
}

// BlackScreeningResponse represents the response from black address screening
type BlackScreeningResponse struct {
	BaseResponse
	Data *BlackScreeningData `json:"data"`
}

// HasAnyRisk checks if the screening result has any risk flags
func (d *BlackScreeningData) HasAnyRisk() bool {
	return d.Sanction || d.Scam || d.Gambling || d.Darknet || d.Theft ||
		d.Mixing || d.Hacker || d.Ransomware || d.Trojan || d.ChildAbuseMaterial ||
		d.Terrorist || d.Drug || d.Lawsuit || d.BusinessBlackList || d.Piracy ||
		d.FraudShop || d.UndergroundBank || d.MoneyMule || d.ProtocolPiracy ||
		d.IllicitActorOrganization || d.HighRiskExchange || d.HighRiskJurisdictionFATF ||
		d.GreyListFATF || d.OfficialFreeze
}
