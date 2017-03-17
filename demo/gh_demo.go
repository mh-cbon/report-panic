// +build ignore

package main

import (
	"github.com/mh-cbon/report-panic"
)

// wrap your main code with reportpanic.HandleMain,
// provide it a reporter such as Gh

func main() {
	reportpanic.HandleMain(reportpanic.Gh("mh-cbon/demotest"), func() {
		panic("oh no!")
	})
}
