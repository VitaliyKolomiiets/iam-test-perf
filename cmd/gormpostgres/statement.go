package gormpostgres

import (
	"fmt"
	"github.com/google/uuid"
	"iam-test-perf/cmd/models"
	"math/rand"
	"strconv"
)

func SearchStatementByParams(principles []string, resources []string, actions []string) {
	var statementIds []int
	NewClient().Raw("select s.id from statements s"+
		"  left join actions a on s.id = a.statement_id "+
		"  left join principles p on s.id = p.statement_id "+
		"  left join resources r on s.id = r.statement_id "+
		"where a.action in (?) and r.krn in (?) and p.krn in (?) ", actions, resources, principles).Scan(&statementIds)

	fmt.Printf("Search applied, result len = %s\n", strconv.Itoa(len(statementIds)))
}

func StatementFilling() {
	client := NewClient()

	for i := 0; i < 100; i++ {
		var theArrayStatement []*models.Statement

		var tenantName = generateRandomString()

		for j := 0; j < 10000; j++ {
			theArrayStatement = append(theArrayStatement, buildStatement(tenantName))
		}

		client.CreateInBatches(theArrayStatement, 1000)
	}

	println("Statement filled")
}

func buildStatement(tenant string) *models.Statement {
	var theArrayAction = []models.Action{{Action: "endpoint:read"}, {Action: "endpoint:write"}, {Action: "endpoint:delete"}}

	var theArrayResource []models.Resource
	for i := 0; i < 10; i++ {
		if i == 0 {
			theArrayResource = append(theArrayResource, models.Resource{Krn: "krn:epr:*"})
		} else if i == 1 {
			theArrayResource = append(theArrayResource, models.Resource{Krn: "krn:epr:" + tenant + "::*"})
		} else {
			theArrayResource = append(theArrayResource, models.Resource{Krn: "krn:epr:" + tenant + "::endpoint/" + uuid.New().String()})
		}
	}

	var theArrayPrinciple []models.Principle
	for i := 0; i < 5; i++ {
		if i == 0 {
			theArrayPrinciple = append(theArrayPrinciple, models.Principle{Krn: "krn:iam:*"})
		} else if i == 1 {
			theArrayPrinciple = append(theArrayPrinciple, models.Principle{Krn: "krn:iam:" + tenant + "::*"})
		} else {
			theArrayPrinciple = append(theArrayPrinciple, models.Principle{Krn: "krn:iam:" + tenant + "::user/" + uuid.New().String()})
		}
	}

	return &models.Statement{Type: "Allow", Actions: theArrayAction, Resources: theArrayResource, Principles: theArrayPrinciple}
}

func generateRandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	s := make([]rune, 10)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
