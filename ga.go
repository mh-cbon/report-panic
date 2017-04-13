package reportpanic

import (
	"fmt"
	"os"
	"strings"

	"github.com/mh-cbon/report-panic/ga"
)

// GaReporter is a reporter to reports starts and panics of your programs to your GA account.
type GaReporter struct {
	ID      string
	program string
	version string
	client  *ga.Client
}

// Ga construct a GA reporter.
func Ga(ID string, program string, version string) *GaReporter {
	c, err := ga.NewClient("go-ga", ID)
	if err != nil {
		panic(err)
	}
	ret := &GaReporter{
		program: program,
		version: version,
		client:  c,
	}
	go func() {
		// <-time.After(time.Second * 2)
		ret.Start()
	}()
	return ret
}

// Start notify program starts to your GA account.
func (g *GaReporter) Start() error {
	if isCI() {
		return nil
	}
	URL := fmt.Sprintf("%v/%v/%v", g.program, g.version, "start")
	return g.client.PageView(URL)
}

// Report notify panics to your GA account.
func (g *GaReporter) Report(p *ParsedPanic) error {
	if isCI() {
		return nil
	}
	URL := fmt.Sprintf("%v/%v/%v", g.program, g.version, "panic")
	return g.client.PageView(URL)
}

func isCI() bool {
	return strings.ToUpper(os.Getenv("CI")) == "TRUE"
}
