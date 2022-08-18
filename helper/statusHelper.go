package helper

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"project/checker"
	"sync"
)

var waitGroup sync.WaitGroup

func GetStatusHelper(websiteList map[string]bool) {
	tbl := table.New(color.YellowString("\t\tWebsite"), color.YellowString("\tStatus"))
	tbl.WithPadding(15)
	for website, _ := range websiteList {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			checkerObj := &checker.StatusChecker{}

			result, _ := checkerObj.CheckStatus(website)
			if result {
				tbl.AddRow("\t"+website, color.GreenString("UP"))
			} else {
				tbl.AddRow("\t"+website, color.RedString("DOWN"))
			}
		}()
		waitGroup.Wait()
	}
	tbl.Print()
}
