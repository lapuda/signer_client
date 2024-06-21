package api

type NewRequest struct{}
type NewResponse struct {
	Address string `json:"address"`
	Pk      string `json:"pk"`
}
