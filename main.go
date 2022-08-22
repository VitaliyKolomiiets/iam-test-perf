package main

import (
	"fmt"
	pg "iam-test-perf/cmd/gormpostgres"
	"iam-test-perf/cmd/models"
	"iam-test-perf/cmd/servce/action"
	"iam-test-perf/cmd/servce/krn"
	"time"
)

func main() {
	/*pg.MigrateEntity()
	pg.FillStatement()*/

	fmt.Println("CASE-1: Evaluate if user has access to one resource")
	/*Function exist, return true or false*/
	actions := action.Action("iam:endpoint:read")
	resourceKRN, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/6fe3bfae-c0bf-4b9e-82c1-57e84e2776f0")
	principleKRN, _ := krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	start := time.Now()
	pg.ExistSearchStatementByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-1: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-2: Evaluate if user has access to several resources.")
	/*Function exist, return true or false*/
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	start = time.Now()
	pg.ExistSearchStatementByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-2: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-3: Retrieve one allowed resource.")
	/*Function select first krn from resource by matched KRNs, return ??*/
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	start = time.Now()
	var resourceKRNs = pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	pg.SearchByKRNsByType(models.ResourcePayload, resourceKRNs)
	fmt.Printf("CASE-3: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-4: Retrieve all resources: all requested resources are allowed.")
	/*Function select krn from resource by matched KRNs, return ??*/
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	start = time.Now()
	resourceKRNs = pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	pg.SearchByKRNsByType(models.ResourcePayload, resourceKRNs)
	fmt.Printf("CASE-4: Search took %s\n", time.Since(start).String())
	/*-----------------------------------------------------------------------------------------------------*/
	fmt.Println("CASE-5: Retrieve top N allowed resources")
	/*Function select krn from resource by matched KRNs, return top??*/
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	fmt.Println("CASE-6: Retrieve N middle resources")
	/*Function select krn from resource by matched KRNs, return middle??*/
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	fmt.Println("CASE-7: Retrieve last N allowed resources")
	/*Function select krn from resource by matched KRNs, return last N??*/
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	fmt.Println("CASE-8: Retrieve one denied resource")
	/*Function select  first krn from resource by matched KRNs and statement = denied, return */
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	fmt.Println("CASE-9: Retrieve all resources: all requested resources are denied")
	/*Function select  first krn from resource by matched KRNs and statement = denied, return */
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	fmt.Println("CASE-10: Retrieve N allowed and N denied resources")
	/*Function select  first krn from resource by matched KRNs and statement = denied, return */
	actions = action.Action("iam:endpoint:read")
	resourceKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::endpoint/*")
	principleKRN, _ = krn.NewKRNFromString("krn:xvlbzgbaic:mrajwwhthc::user/bdb2771d-e62c-417d-af52-e61bf197aff0")

	start = time.Now()
	pg.SearchStatementIdsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-1: Search took %s\n", time.Since(start).String())

	start = time.Now()
	pg.ExistSearchStatementByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-2: Search took %s\n", time.Since(start).String())

	start = time.Now()
	pg.SearchPrincipleKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-3: Search took %s\n", time.Since(start).String())

	start = time.Now()
	resourceKRNs = pg.SearchResourceKRNsByParams(principleKRN.MatchingKRNs(), resourceKRN.MatchingKRNs(), actions.MatchingActionsString())
	fmt.Printf("CASE-4: Search took %s\n", time.Since(start).String())

	start = time.Now()
	pg.SearchByKRNsByType(models.ResourcePayload, resourceKRNs)
	fmt.Printf("CASE-5: Search took %s\n", time.Since(start).String())
}
