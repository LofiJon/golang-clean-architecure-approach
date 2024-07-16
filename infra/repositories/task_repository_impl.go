package repositories

import (
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/core/repositories"
	_ "golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/infra/entities"
	"gorm.io/gorm"
	"time"
)

var _ repositories.TaskRepository = &TaskRepositoryImpl{}

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db}
}

func (r *TaskRepositoryImpl) Create(task *models.Task) error {
	taskEntity := entities.ToTaskEntity(task)
	return r.db.Create(taskEntity).Error
}

func (r *TaskRepositoryImpl) GetAll() ([]models.Task, error) {
	var taskEntities []entities.TaskEntity
	err := r.db.Find(&taskEntities).Error
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, taskEntity := range taskEntities {
		tasks = append(tasks, *taskEntity.ToCoreTask())
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) GetByID(id string) (models.Task, error) {
	var taskEntity entities.TaskEntity
	err := r.db.First(&taskEntity, "id = ?", id).Error
	if err != nil {
		return models.Task{}, err
	}
	return *taskEntity.ToCoreTask(), nil
}

func (r *TaskRepositoryImpl) Update(task *models.Task) error {
	taskEntity := entities.ToTaskEntity(task)
	return r.db.Save(taskEntity).Error
}

func (r *TaskRepositoryImpl) Delete(id string) error {
	var taskEntity entities.TaskEntity
	if err := r.db.First(&taskEntity, "id = ?", id).Error; err != nil {
		return err
	}

	taskEntity.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	return r.db.Save(&taskEntity).Error
}

func (r *TaskRepositoryImpl) Count() (int64, error) {
	var count int64
	err := r.db.Model(&entities.TaskEntity{}).Count(&count).Error
	return count, err
}

func (r *TaskRepositoryImpl) GetPaged(page, perPage int) ([]models.Task, error) {
	var taskEntities []entities.TaskEntity
	offset := (page - 1) * perPage
	err := r.db.Where("deleted_at IS NULL").Limit(perPage).Offset(offset).Find(&taskEntities).Error
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, entity := range taskEntities {
		tasks = append(tasks, *entity.ToCoreTask())
	}

	return tasks, nil
}
