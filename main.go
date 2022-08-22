package main

import (
	"fmt"
	pg "iam-test-perf/cmd/gormpostgres"
	"iam-test-perf/cmd/servce/action"
	"iam-test-perf/cmd/servce/krn"
	"time"
)

func main() {
	/*pg.MigrateEntity()
	pg.FillStatement()*/

	fmt.Println("CASE-1: Evaluate if user has access to one resource")
	actions := action.Action("iam:endpoint:read")
	resourceKRN, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/f78c7a89-a46a-4b53-8071-0fd1728b0648")
	principleKRN, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/dd4a7473-97cd-4911-b758-360efbe45a85")

	start := time.Now()
	pg.ExistSearchStatementByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-1: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-2: Evaluate if user has access to two resources.")
	actions = action.Action("iam:endpoint:read")
	resourcesKRN1, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/422aee60-f367-44a7-b26a-55c3cdaf26fc")
	resourcesKRN2, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/f78c7a89-a46a-4b53-8071-0fd1728b0648")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/dd4a7473-97cd-4911-b758-360efbe45a85")

	start = time.Now()
	pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), append(resourcesKRN1.MatchingKRNs(), resourcesKRN2.MatchingKRNs()...), actions.MatchingActionsString())
	fmt.Printf("CASE-2: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-3: Evaluate if user has access to several resources.")
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/dd4a7473-97cd-4911-b758-360efbe45a85")

	start = time.Now()
	pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-3: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	//TODO fix
	fmt.Println("CASE-4: Retrieve one allowed resource.")
	pg.FillStatementOneAllowedResource()

	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:iam:kaa::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:iam:kaa::user/11111111-8c1c-45e9-a137-c8ed88d2a722")

	start = time.Now()
	pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-4: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-5: Retrieve all resources: all requested resources are allowed.")
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:byqvpcufqp:*")
	principleKRN, _ = krn.NewKRNFromString("krn:byqvpcufqp:rzuhtifbzd::*")

	start = time.Now()
	pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-5: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-6: Retrieve all resources: all requested resources are allowed.")
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:byqvpcufqp:obdayfnyyh::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:byqvpcufqp:*")

	start = time.Now()
	pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-6: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-7: Retrieve one grouped by state")
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:byqvpcufqp:*")
	principleKRN, _ = krn.NewKRNFromString("krn:byqvpcufqp:*")

	start = time.Now()
	pg.SearchResourceKRNsByParamsAndGroupedByState(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-7: Search took %s\n", time.Since(start).String())
}
