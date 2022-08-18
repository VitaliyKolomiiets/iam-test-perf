package main

import (
	"fmt"
	pg "iam-test-perf/cmd/gormpostgres"
	"iam-test-perf/cmd/servce/action"
	"iam-test-perf/cmd/servce/krn"
	"time"
)

func main() {
	//pg.StatementMigration()

	actions := action.Action("iam:user:endpoint:read")
	resourceKRN, _ := krn.NewKRNFromString("krn:iam:kaa::endpoint/5766b7e9-1f16-443d-8e4a-553f70733aa7")
	principleKRN, _ := krn.NewKRNFromString("krn:iam:kaa::user/vitaliy")

	start := time.Now()

	pg.ExecuteSearch(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())

	fmt.Printf("Search took %s", time.Since(start).String())
}
