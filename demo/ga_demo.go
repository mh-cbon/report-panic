// +build ignore

package main

import (
	"github.com/mh-cbon/report-panic"
)

// wrap your main code with reportpanic.HandleMain,
// provide it a reporter such as Gh

func main() {
	// Don t forget to change the GA ID.
	reportpanic.HandleMain(reportpanic.Ga("UA-93911415-1", "my-program", "0.0.1"), func() {
		panic("oh no!")
	})
}
