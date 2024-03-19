package model

type Payload struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestID   string `json:"requestId"`
	Amount      string `json:"amount"`
	OrderID     string `json:"orderId"`
	OrderInfo   string `json:"orderInfo"`
	RedirectUrl string `json:"redirectUrl"`
	IpnUrl      string `json:"ipnUrl"`
	ExtraData   string `json:"extraData"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
}

type PosHash struct {
	PartnerCode  string `json:"partnerCode"`
	PartnerRefID string `json:"partnerRefId"`
	Amount       int    `json:"amount"`
	PaymentCode  string `json:"paymentCode"`
}

type PosPayload struct {
	PartnerCode  string `json:"partnerCode"`
	PartnerRefID string `json:"partnerRefId"`
	Hash         string `json:"hash"`
	Version      int    `json:"version"`
}
