package gormpostgres

import (
	"fmt"
	"github.com/google/uuid"
	"iam-test-perf/cmd/models"
	"math/rand"
	"strconv"
)

type result struct {
	Type string
	list []string
}

func SearchStatementIdsByParams(principles []string, resources []string, actions []string) {
	var statementIds []int
	NewClient().Raw("select distinct s.id from statements s"+
		"  left join actions a on s.id = a.statement_id "+
		"  left join principles p on s.id = p.statement_id "+
		"  left join resources r on s.id = r.statement_id "+
		"where a.action in (?) and r.krn in (?) and p.krn in (?) ", actions, resources, principles).Scan(&statementIds)

	fmt.Printf("Search applied, result len = %s\n", strconv.Itoa(len(statementIds)))
}

func ExistSearchStatementByParams(principles []string, resources []string, actions []string) {
	var exist bool
	NewClient().Raw("select exists(select s.id from statements s"+
		"  left join actions a on s.id = a.statement_id "+
		"  left join principles p on s.id = p.statement_id "+
		"  left join resources r on s.id = r.statement_id "+
		"where a.action in (?) and r.krn in (?) and p.krn in (?)) ", actions, resources, principles).Scan(&exist)

	fmt.Println("Search applied, result: =" + strconv.FormatBool(exist))
}

func SearchPrincipleKRNsByParams(principles []string, resources []string, actions []string) {
	var principleKRNs []string
	NewClient().Raw("select distinct p.krn from statements s"+
		"  left join actions a on s.id = a.statement_id "+
		"  left join principles p on s.id = p.statement_id "+
		"  left join resources r on s.id = r.statement_id "+
		"where a.action in (?) and r.krn in (?) and p.krn in (?) ", actions, resources, principles).Scan(&principleKRNs)

	fmt.Println("Search applied, result: =" + strconv.Itoa(len(principleKRNs)))
}

func SearchResourceKRNsByParams(principles []string, resources []string, actions []string) []string {
	var resourceKRNs []string
	NewClient().Raw("select distinct r.krn from statements s"+
		"  left join actions a on s.id = a.statement_id "+
		"  left join principles p on s.id = p.statement_id "+
		"  left join resources r on s.id = r.statement_id "+
		"where a.action in (?) and r.krn in (?) and p.krn in (?) ", actions, resources, principles).Scan(&resourceKRNs)

	fmt.Println("Search applied, result: =" + strconv.Itoa(len(resourceKRNs)))
	return resourceKRNs
}

func SearchResourceKRNsByParamsAndGroupedByState(principles []string, resources []string, actions []string) {
	var resourceKRNs []result
	NewClient().Raw("select  s.type, array_agg(distinct r.krn) as list from statements s"+
		"  left join actions a on s.id = a.statement_id "+
		"  left join principles p on s.id = p.statement_id "+
		"  left join resources r on s.id = r.statement_id "+
		"where a.action in (?) and r.krn in (?) and p.krn in (?) "+
		"group by s.type ", actions, resources, principles).Scan(&resourceKRNs)

	fmt.Println("Search applied, result: =" + strconv.Itoa(len(resourceKRNs)))
}

func FillStatementOneAllowedResource() {
	var theArrayAction = []models.Action{{Action: "iam:endpoint:read"}}
	var theArrayResource = []models.Resource{{Krn: "krn:iam:kaa::endpoint/11111111-f367-44a7-b26a-55c3cdaf26fc"}}
	var theArrayPrinciple = []models.Principle{{Krn: "krn:iam:kaa::user/11111111-8c1c-45e9-a137-c8ed88d2a722"}}

	NewClient().Create(&models.Statement{Type: "Allow", Actions: theArrayAction, Principles: theArrayPrinciple, Resources: theArrayResource})
}

func FillStatement() {
	client := NewClient()

	for i := 0; i < 100; i++ {
		var theArrayStatement []*models.Statement

		serviceName := generateRandomString()

		for j := 0; j < 10000; j++ {
			if j%10 == 0 {
				theArrayStatement = append(theArrayStatement, buildStatement(serviceName, generateRandomString(), true, i > 80))
			} else {
				theArrayStatement = append(theArrayStatement, buildStatement(serviceName, generateRandomString(), false, i > 80))
			}
		}

		client.CreateInBatches(theArrayStatement, 1000)
	}

	println("Statement filled")
}

func buildStatement(serviceName string, tenantName string, includeServiceWildcard bool, isStateAllow bool) *models.Statement {
	var theArrayAction = []models.Action{{Action: "iam:endpoint:read"}, {Action: "iam:endpoint:write"}, {Action: "iam:endpoint:delete"}}

	var theArrayResource []models.Resource
	for i := 0; i < 10; i++ {
		if i == 0 && includeServiceWildcard {
			theArrayResource = append(theArrayResource, models.Resource{Krn: "krn:" + serviceName + ":*"})
		} else if i == 1 {
			theArrayResource = append(theArrayResource, models.Resource{Krn: "krn:" + serviceName + ":" + tenantName + "::*"})
		} else {
			theArrayResource = append(theArrayResource, models.Resource{Krn: "krn:" + serviceName + ":" + tenantName + "::endpoint/" + uuid.New().String()})
		}
	}

	var theArrayPrinciple []models.Principle
	for i := 0; i < 5; i++ {
		if i == 0 && includeServiceWildcard {
			theArrayPrinciple = append(theArrayPrinciple, models.Principle{Krn: "krn:" + serviceName + ":*"})
		} else if i == 1 {
			theArrayPrinciple = append(theArrayPrinciple, models.Principle{Krn: "krn:" + serviceName + ":" + tenantName + "::*"})
		} else {
			theArrayPrinciple = append(theArrayPrinciple, models.Principle{Krn: "krn:" + serviceName + ":" + tenantName + "::user/" + uuid.New().String()})
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
