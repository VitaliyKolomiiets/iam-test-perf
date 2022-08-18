package gormpostgres

import (
	"encoding/base64"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"iam-test-perf/cmd/models"
	"math/rand"
)

func StatementMigration() {
	dsn := "host=localhost user=iam1 password=root1 dbname=iam1 port=5433"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err1 := db.AutoMigrate(&models.Statement{}, &models.Principle{}, &models.Action{}, &models.Resource{})
	if err1 != nil {
		panic("Error happened during migration")
	}

	for i := 0; i < 100; i++ {
		var theArrayStatement []*models.Statement

		for i := 0; i < 10000; i++ {
			theArrayStatement = append(theArrayStatement, buildStatement())
		}

		db.CreateInBatches(theArrayStatement, 1000)
	}

	println("done")
}

func buildStatement() *models.Statement {
	var theArrayAction = []models.Action{{Action: "endpoint:read"}, {Action: "endpoint:write"}, {Action: "endpoint:delete"}}

	var theArrayResource []models.Resource
	for i := 0; i < 10; i++ {
		theArrayResource = append(theArrayResource, models.Resource{Krn: "krn:epr:3fdfde93-661a-47d1-abcc-68452dd320c7::endpoint/" + generateRandomString()})
	}

	var theArrayPrinciple []models.Principle
	for i := 0; i < 5; i++ {
		theArrayPrinciple = append(theArrayPrinciple, models.Principle{Krn: "krn:epr:3fdfde93-661a-47d1-abcc-68452dd320c7::endpoint/" + generateRandomString()})
	}

	return &models.Statement{Type: "Allow", Actions: theArrayAction, Resources: theArrayResource, Principles: theArrayPrinciple}
}

func generateRandomString() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)

	return base64.StdEncoding.EncodeToString(b)
}
