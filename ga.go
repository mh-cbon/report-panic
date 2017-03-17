package reportpanic

import (
	"fmt"
	"os"
	"strings"

	"github.com/jpillora/go-ogle-analytics"
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
	c, err := ga.NewClient(ID)
	if err != nil {
		panic(err)
	}
	ret := &GaReporter{
		program: program,
		version: version,
		client:  c,
	}
	go ret.Start()
	return ret
}

// Start notify program starts to your GA account.
func (g *GaReporter) Start() error {
	if strings.ToUpper(os.Getenv("CI")) == "TRUE" {
		return nil // skip CI environments.
	}
	URL := fmt.Sprintf("http://%v/%v/%v/%v", "localhost", g.program, g.version, "start")
	return g.client.Send(ga.NewPageview(URL))
	// return g.client.Send(&myPV{url: URL})
	// return g.client.Send(ga.NewEvent(g.Category, "start").Label(g.Label))
}

// Report notify panics to your GA account.
func (g *GaReporter) Report(p *ParsedPanic) error {
	if strings.ToUpper(os.Getenv("CI")) == "TRUE" {
		return nil // skip CI environments.
	}
	URL := fmt.Sprintf("http://%v/%v/%v/%v", "localhost", g.program, g.version, "start")
	return g.client.Send(ga.NewPageview(URL))
	// return g.client.Send(&myPV{url: URL})
}
