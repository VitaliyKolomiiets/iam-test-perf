package main

import (
	"fmt"
	pg "iam-test-perf/cmd/gormpostgres"
	"iam-test-perf/cmd/servce/action"
	"iam-test-perf/cmd/servce/krn"
	"time"
)

func main() {
	pg.MigrateEntity()
	pg.FillStatement()

	actions := action.Action("iam:endpoint:read")
	resourceKRN, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/6fe3bfae-c0bf-4b9e-82c1-57e84e2776f0")
	principleKRN, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	start := time.Now()
	pg.SearchStatementIdsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-1: Search took %s\n", time.Since(start).String())

	start = time.Now()
	pg.ExistSearchStatementByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-2: Search took %s\n", time.Since(start).String())

	start = time.Now()
	pg.SearchPrincipleKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-3: Search took %s\n", time.Since(start).String())

	start = time.Now()
	pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-4: Search took %s\n", time.Since(start).String())
}
