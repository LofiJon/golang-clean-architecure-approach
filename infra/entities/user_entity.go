package entities

import (
	"github.com/google/uuid"
	"golang-api-clean-architecture/core/models"
	"gorm.io/gorm"
	"time"
)

type UserEntity struct {
	BaseEntity
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (UserEntity) TableName() string {
	return "users"
}

// Mapping user to user entity
func ToUserEntity(user *models.User) *UserEntity {
	id, _ := uuid.Parse(user.ID)
	var deletedAt gorm.DeletedAt
	if user.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *user.DeletedAt,
			Valid: true,
		}
	}

	return &UserEntity{
		BaseEntity: BaseEntity{
			ID:        id,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: deletedAt,
		},
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Username: user.Username,
	}
}

// Mapping user entity to user
func (user *UserEntity) ToCoreUser() *models.User {
	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}
	return &models.User{
		BaseModel: models.BaseModel{
			ID:        user.ID.String(),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: deletedAt,
		},
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Username: user.Username,
	}
}
