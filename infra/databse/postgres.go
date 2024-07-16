package database

import (
	"golang-api-clean-architecture/infra/entities"
	"golang-api-clean-architecture/infra/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=example_user password=example_password dbname=go_basic port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Add extension for UUID if not exists
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.AutoMigrate(&entities.TaskEntity{})
	seeds.TaskSeed(db)

	log.Println("Database migrated")

	return db
}
