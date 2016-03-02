package http

type registerNodeRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type registerNodeResponse struct {
	Status string `json:"status"`
}
