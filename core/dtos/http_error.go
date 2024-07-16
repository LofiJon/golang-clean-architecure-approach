package dtos

// HTTPError represents an HTTP error response
// swagger:model HTTPError
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
