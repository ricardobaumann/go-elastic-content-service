package content

type contentRequest struct {
	ID string `json:"id"`
}

type contentResponse struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

type contentInput struct {
	ID   string
	Body string
}
