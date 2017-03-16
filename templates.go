package reportpanic

import (
	"bytes"
	"runtime"
	"text/template"
)

// StdTemplateResolver is the default templater for github reports.
type StdTemplateResolver struct {
	GoVersion string
	Os        string
	Arch      string
	Reason    string
	Panic     string
	Template  string
}

// NewStdTemplateResolver is a ctor.
func NewStdTemplateResolver(tpl string) *StdTemplateResolver {
	return &StdTemplateResolver{
		Template: tpl,
	}
}

// Make a report content to post in the github issue.
func (c *StdTemplateResolver) Make(p *ParsedPanic) (string, error) {
	c.GoVersion = runtime.Version()
	c.Os = runtime.GOOS
	c.Arch = runtime.GOARCH
	c.Panic = "```\n" + p.Content + "\n```" // ugly, but workable.
	c.Reason = p.Reason

	var err error
	t := template.New("some")
	t, err = t.Parse(c.Template)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = t.Execute(&b, c)
	return b.String(), err
}

// PanicReportTemplates contains templates to build the body report.
var PanicReportTemplates = map[string]string{
	"md": `
  __Go__: {{.GoVersion}}

  __OS__: {{.Os}}

  __Arch__: {{.Arch}}

  {{.Panic}}
  `,
}
