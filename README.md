# report-panic


[![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/report-panic)](https://goreportcard.com/report/github.com/mh-cbon/report-panic)

[![GoDoc](https://godoc.org/github.com/mh-cbon/report-panic?status.svg)](http://godoc.org/github.com/mh-cbon/report-panic)


Package reportpanic Automatically report your programs panic.


# API example

### Report to a github repository


###### > demo/gh_demo.go
```go
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
```

### Report to a google analytics account


###### > demo/ga_demo.go
```go
// +build ignore

package main

import (
	"github.com/mh-cbon/report-panic"
)

// wrap your main code with reportpanic.HandleMain,
// provide it a reporter such as Gh

func main() {
	reportpanic.HandleMain(reportpanic.Ga("UA-93911415-1", "my-program", "0.0.1"), func() {
		panic("oh no!")
	})
}
```

# Install

#### Glide

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon/report-panic
cd $GOPATH/src/github.com/mh-cbon/report-panic
git clone https://github.com/mh-cbon/report-panic.git .
glide install
go install
```


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

# Credits

[Mitchell Hashimoto](https://github.com/mitchellh) for the base code provided in [panicwrap](https://github.com/mitchellh/panicwrap)
