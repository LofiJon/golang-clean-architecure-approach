package seeds

import (
	"github.com/google/uuid"
	"golang-api-clean-architecture/infra/entities"
	"gorm.io/gorm"
	"log"
	"time"
)

func TaskSeed(db *gorm.DB) {
	// Seeding initials tasks to show
	tasks := []entities.TaskEntity{
		{
			BaseEntity: entities.BaseEntity{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "First Task",
			Done: false,
		},
		{
			BaseEntity: entities.BaseEntity{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Second Task",
			Done: true,
		},
		{
			BaseEntity: entities.BaseEntity{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Third Task",
			Done: true,
		},
		{
			BaseEntity: entities.BaseEntity{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Fourth Task",
			Done: true,
		},
	}

	// Insert the seed data into the database
	for _, task := range tasks {
		if err := db.Create(&task).Error; err != nil {
			log.Fatalf("Could not seed tasks: %v", err)
		}
	}
}
