package api

type ListRequest struct{}

type ListResponse struct {
	Items []string `json:"items"`
}
