package api

type ImportRequest struct {
	Pk string `json:"pk"`
}
type ImportResponse struct {
	Code    int    `json:"code"`
	Address string `json:"address"`
}
