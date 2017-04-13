# {{.Name}}

{{template "badge/goreport" .}}{{template "badge/godoc" .}}

{{pkgdoc "index.go"}}

# {{toc 5}}

# API example

### Report to a github repository

{{cat "demo/gh_demo.go" | color "go"}}

### Report to a google analytics account

{{cat "demo/ga_demo.go" | color "go"}}

# Install

#### Glide
{{template "glide/install" .}}

# Ga Configuration

You can only set
- GA ID
- Your program name
- Your program version

The GA reporter will emit `pageview` on program `start`, `panic` as
`http://localhost/<program>/<version>/<action>`

# Gh Configuration

### Change the github token reporter

To set a different token to report panics,
add this flag to your build compilation
`--ldflags "-X reportpanic.GhReporterToken=xxxxx"`


```sh
go build --ldflags "-X reportpanic.GhReporterToken=xxxxx"
```

### Change the github template generating the body report

1. Access to the templater instance of the reporter

```go
package main

import (
  "github.com/mh-cbon/report-panic"
)

func main () {
  reporter := reportpanic.Gh("your_user/your_repo")
  reporter.Templater.Template = "template { {content} } here"
  reportpanic.HandleMain(reporter, func () {
    panic("oh no!")
  })
}
```

2. Or, use ldflags at build time

At build time, use `ldflags` to configure the template

```sh
go build --ldflags "-X reportpanic.GhTemplate=your_tpl"
```

Register `your_tpl` to panicreport

```go
package main

import (
  "github.com/mh-cbon/report-panic"
)

// register your template in the init, or main func
func init(){
  panicreport.PanicReportTemplates["your_tpl"] = "template { {content} } here"
}

func main() {
  reportpanic.HandleMain(reportpanic.Gh("your_user/your_repo"), func () {
    panic("oh no!")
  })
}
```

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)

# Credits

[Mitchell Hashimoto](https://github.com/mitchellh) for the base code provided in [panicwrap](https://github.com/mitchellh/panicwrap)
