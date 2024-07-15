package dtos

type PageableDto struct {
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	Items      interface{} `json:"items"`
}
