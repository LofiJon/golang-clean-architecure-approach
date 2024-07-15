package models

type Task struct {
	BaseModel
	Name string `json:"name"`
	Done bool   `json:"done"`
}
