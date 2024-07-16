package seeds

import (
	"golang-api-clean-architecture/infra/entities"
	"gorm.io/gorm"
	"log"
)

func TaskSeed(db *gorm.DB) {
	// Seeding initials tasks to show
	tasks := []entities.TaskEntity{
		{

			Name: "First Task",
			Done: false,
		},
		{

			Name: "Second Task",
			Done: true,
		},
		{

			Name: "Third Task",
			Done: true,
		},
		{

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
