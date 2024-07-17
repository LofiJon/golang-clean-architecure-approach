package entities

import (
	"golang-api-clean-architecture/core/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name"`
	Done      bool           `json:"done"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (e *TaskEntity) ToCoreTask() *models.Task {
	var deletedAt *time.Time
	if e.DeletedAt.Valid {
		deletedAt = &e.DeletedAt.Time
	}
	return &models.Task{
		BaseModel: models.BaseModel{
			ID:        e.ID.String(),
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
			DeletedAt: deletedAt,
		},
		Name: e.Name,
		Done: e.Done,
	}
}

func ToTaskEntity(task *models.Task) *TaskEntity {
	if task == nil {
		return nil
	}

	id, _ := uuid.Parse(task.ID)
	var deletedAt gorm.DeletedAt
	if task.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *task.DeletedAt,
			Valid: task.DeletedAt != nil,
		}
	}
	return &TaskEntity{
		ID:        id,
		Name:      task.Name,
		Done:      task.Done,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
