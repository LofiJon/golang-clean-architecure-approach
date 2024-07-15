package models

import "time"

type BaseModel struct {
	ID        string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
