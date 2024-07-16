package repositories

import (
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/infra/entities"
	"gorm.io/gorm"
	"time"
)

var _ repositories.UserRepository = &UserRepositoryImpl{}

type UserRepositoryImpl struct {
	database *gorm.DB
}

func (u UserRepositoryImpl) Create(user *models.User) error {
	//TODO implementing create user
	userEntity := entities.ToUserEntity(user)
	return u.database.Create(userEntity).Error
}

func (u UserRepositoryImpl) GetAll() ([]models.User, error) {
	//TODO implement get all users
	var userEntities []entities.UserEntity
	err := u.database.Find(&userEntities).Error
	if err != nil {
		return nil, err
	}

	var users []models.User
	for _, userEntity := range userEntities {
		users = append(users, *userEntity.ToCoreUser())
	}
	return users, nil
}

func (u UserRepositoryImpl) GetByID(id string) (models.User, error) {
	//TODO implement get user by id
	var userEntity entities.UserEntity
	err := u.database.First(&userEntity, "id = ?", id).Error
	if err != nil {
		return models.User{}, err
	}
	return *userEntity.ToCoreUser(), nil
}

func (u UserRepositoryImpl) Update(user *models.User) error {
	//TODO implement update user
	userEntity := entities.ToUserEntity(user)
	return u.database.Save(userEntity).Error
}

func (u UserRepositoryImpl) Delete(id string) error {
	//TODO implement soft delete user by id
	var userEntity entities.UserEntity
	if err := u.database.First(&userEntity, "id = ?", id).Error; err != nil {
		return err
	}

	userEntity.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	return u.database.Save(&userEntity).Error
}

func (u UserRepositoryImpl) Count() (int64, error) {
	//TODO implement count users
	var count int64
	err := u.database.Model(&entities.UserEntity{}).Count(&count).Error
	return count, err
}

func (u UserRepositoryImpl) GetPaged(page, size int) ([]models.User, error) {
	//TODO implement page users
	var userEntities []entities.UserEntity
	offset := (page - 1) * size
	err := u.database.Where("deleted_at IS NULL").Limit(size).Offset(offset).Find(&userEntities).Error
	if err != nil {
		return nil, err
	}

	var users []models.User
	for _, entity := range userEntities {
		users = append(users, *entity.ToCoreUser())
	}

	return users, nil
}
