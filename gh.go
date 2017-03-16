package reportpanic

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GhReporterToken is a gh token to be able to post issues
var GhReporterToken = "b" + "31198d95cb24e3e2ae6c3f11141df1d417834eb"

// GhTemplate is the name of the template to generate the report
var GhTemplate = "md"

// GhReporter is a reporter to reports panics of your programs to their github repository.
type GhReporter struct {
	Token      string
	Owner      string
	Repo       string
	IssueTitle string
	Templater  TemplateResolver
}

// Gh is the default constructor for a GhReporter.
func Gh(repo string) *GhReporter {
	ret := GhReporter{}
	if repos := strings.Split(repo, "/"); len(repos) == 2 {
		ret.Owner = repos[0]
		ret.Repo = repos[1]
	} else {
		panic("repo must be a string like owner/name")
	}
	// this is a built-in token to create issues as report-panic-bot.
	ret.Token = GhReporterToken
	ret.IssueTitle = "Automatic panic report"
	tpl := PanicReportTemplates[GhTemplate]
	ret.Templater = NewStdTemplateResolver(tpl)
	return &ret
}

// Report create github issues.
func (g *GhReporter) Report(p *ParsedPanic) error {

	body, err := g.Templater.Make(p)
	if err != nil {
		return err
	}

	title := g.IssueTitle
	if p.Reason != "" {
		title = p.Reason
	}

	input := &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	// TDB, avoid creation of multiple issues for the same error.
	// need to interpret the panic content,
	// so the program can identify similar/different panics.
	_, _, err = client.Issues.Create(context.Background(), g.Owner, g.Repo, input)

	return err
}
