package repositories

import (
	"errors"
	"golang-api-clean-architecture/core/models"
	repository "golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/infra/entities"
	"gorm.io/gorm"
	"time"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

// Ensure TaskRepositoryImpl implements the TaskRepository interface
var _ repository.TaskRepository = &TaskRepositoryImpl{}

// NewTaskRepository creates a new TaskRepositoryImpl
func NewTaskRepository(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db}
}

func (r *TaskRepositoryImpl) Create(task *models.Task) error {
	taskEntity := entities.ToTaskEntity(task)
	return r.db.Create(taskEntity).Error
}

func (r *TaskRepositoryImpl) GetByID(id string) (models.Task, error) {
	var taskEntity entities.TaskEntity
	err := r.db.First(&taskEntity, "id = ?", id).Error
	if err != nil {
		return models.Task{}, err
	}
	return *taskEntity.ToCoreTask(), nil
}

func (r *TaskRepositoryImpl) GetAll() ([]models.Task, error) {
	var taskEntities []entities.TaskEntity
	err := r.db.Find(&taskEntities).Error
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, entity := range taskEntities {
		tasks = append(tasks, *entity.ToCoreTask())
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) Update(task *models.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	taskEntity := entities.ToTaskEntity(task)
	if taskEntity == nil {
		return errors.New("failed to convert task to entity")
	}
	return r.db.Save(taskEntity).Error
}
func (r *TaskRepositoryImpl) Delete(id string) error {
	var taskEntity entities.TaskEntity
	if err := r.db.First(&taskEntity, "id = ?", id).Error; err != nil {
		return err
	}

	taskEntity.DeletedAt.Valid = true
	taskEntity.DeletedAt.Time = time.Now()
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
