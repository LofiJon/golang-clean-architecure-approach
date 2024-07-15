package requests

type TaskRequest struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}
