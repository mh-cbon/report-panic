package reportpanic

import (
	"fmt"
	"strings"
)

// ParsedPanic is the result of the parsing of a panic output.
type ParsedPanic struct {
	Reason  string
	Content string
}

// String returns a string representation
func (p ParsedPanic) String() string {
	return fmt.Sprintf("%v\n%v", p.Reason, p.Content)
}

// ParsePanic parses a panic output
func ParsePanic(panicContent string) *ParsedPanic {
	reason := ""
	prefix := "panic: "
	j := len(prefix)
	if len(panicContent) >= j && panicContent[0:j] == prefix {
		lines := strings.Split(panicContent, "\n")
		if len(lines) > 0 {
			reason = lines[0][j:]
		}
	}

	ret := ParsedPanic{}
	ret.Content = panicContent
	ret.Reason = reason

	return &ret
}
