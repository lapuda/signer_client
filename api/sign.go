package api

type SignRequest struct {
	SignerAddress string `json:"signerAddress"`
	MessageHash   string `json:"hash"`
}
type SignResponse struct {
	Signature string `json:"signature"`
}
