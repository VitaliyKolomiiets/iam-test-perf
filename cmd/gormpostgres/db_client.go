package gormpostgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"iam-test-perf/cmd/models"
)

func NewClient() *gorm.DB {
	dsn := "host=localhost user=iam1 password=root1 dbname=iam1 port=5433"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func MigrateEntity() {
	err := NewClient().AutoMigrate(&models.Statement{}, &models.Principle{}, &models.Action{}, &models.Resource{}, &models.User{})
	if err != nil {
		panic("Error happened during migration")
	}
}
