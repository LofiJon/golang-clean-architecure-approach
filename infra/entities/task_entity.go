package entities

import (
	"github.com/google/uuid"
	"golang-api-clean-architecture/core/models"
	"gorm.io/gorm"
	"time"
)

type TaskEntity struct {
	BaseEntity
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// TableName sets the name of the table in the database.
func (TaskEntity) TableName() string {
	return "tasks"
}

func ToTaskEntity(task *models.Task) *TaskEntity {
	id, _ := uuid.Parse(task.ID)
	var deletedAt gorm.DeletedAt
	if task.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *task.DeletedAt,
			Valid: true,
		}
	}

	return &TaskEntity{
		BaseEntity: BaseEntity{
			ID:        id,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
			DeletedAt: deletedAt,
		},
		Name: task.Name,
		Done: task.Done,
	}
}

func (t *TaskEntity) ToCoreTask() *models.Task {
	var deletedAt *time.Time
	if t.DeletedAt.Valid {
		deletedAt = &t.DeletedAt.Time
	}
	return &models.Task{
		BaseModel: models.BaseModel{
			ID:        t.ID.String(),
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			DeletedAt: deletedAt,
		},
		Name: t.Name,
		Done: t.Done,
	}
}
