package repositories

import (
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/infra/entities"
	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(database *gorm.DB) *TaskRepository {
	return &TaskRepository{database}
}

func (r *TaskRepository) Create(task *models.Task) error {
	taskEntity := entities.ToTaskEntity(task)
	return r.database.Create(taskEntity).Error
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var taskEntities []entities.TaskEntity
	err := r.database.Find(&taskEntities).Error
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, taskEntity := range taskEntities {
		tasks = append(tasks, *taskEntity.ToCoreTask())
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(id string) (models.Task, error) {
	var taskEntity entities.TaskEntity
	err := r.database.First(&taskEntity, "id = ?", id).Error
	if err != nil {
		return models.Task{}, err
	}
	return *taskEntity.ToCoreTask(), nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	taskEntity := entities.ToTaskEntity(task)
	return r.database.Save(taskEntity).Error
}

func (r *TaskRepository) Delete(id string) error {
	return r.database.Delete(&entities.TaskEntity{}, "id = ?", id).Error
}

func (r *TaskRepository) Count() (int64, error) {
	var count int64
	err := r.database.Model(&entities.TaskEntity{}).Count(&count).Error
	return count, err
}

func (r *TaskRepository) GetPaged(page, size int) ([]models.Task, error) {
	var taskEntities []entities.TaskEntity
	offset := (page - 1) * size
	err := r.database.Limit(size).Offset(offset).Find(&taskEntities).Error
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, taskEntity := range taskEntities {
		tasks = append(tasks, *taskEntity.ToCoreTask())
	}
	return tasks, nil
}
