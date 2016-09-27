// +build ignore

package main

import (
	"github.com/mh-cbon/report-panic"
)

var MyRepo = "mh-cbon/demotest"

func main() {
	reportpanic.HandleMain(reportpanic.Gh(MyRepo), func() {
		panic("oh no!")
	})
}
