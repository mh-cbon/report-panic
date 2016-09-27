package reportpanic

import (
	"os"

	"github.com/mitchellh/panicwrap"
)

// Wraps your main to handle panics.
func HandleMain(r Reporter, some func()) error {
	exitStatus, err := panicwrap.BasicWrap(handlePanic(r))
	if err != nil {
		return err
	}

	if exitStatus >= 0 {
		// If exitStatus >= 0, then we're the parent process and the panicwrap
		// re-executed ourselves and completed. Just exit with the proper status.
		os.Exit(exitStatus)
	}

	// Otherwise, exitStatus < 0 means we're the child.
	// Continue executing your program.
	some()

	return nil
}

// Handles panic reporting.
func handlePanic(r Reporter) panicwrap.HandlerFunc {
	return func(panicContent string) {
		p := ParsePanic(panicContent)
		err := r.Report(p)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

// Reporter must implement a method to Report the panic.
type Reporter interface {
	Report(p *ParsedPanic) error
}

// TemplateResolver must implement a method to Make the report body.
type TemplateResolver interface {
	Make(p *ParsedPanic) (string, error)
}
