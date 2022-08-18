package gormpostgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
)

func ExecuteSearch(principles []string, resources []string, actions []string) {
	dsn := "host=localhost user=iam1 password=root1 dbname=iam1 port=5433 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect database")
	}

	var statementIds []int
	db.Raw("select s.id from statements s"+
		"  left join actions a on s.id = a.statement_id "+
		"  left join principles p on s.id = p.statement_id "+
		"  left join resources r on s.id = r.statement_id "+
		"where a.action in (?) and r.krn in (?) and p.krn in (?) ", actions, resources, principles).Scan(&statementIds)

	fmt.Printf("Search applied, result len = %s\n", strconv.Itoa(len(statementIds)))
}
