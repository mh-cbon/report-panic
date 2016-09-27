# report-panic

Automatically report your programs panic to their github repository.

Simply wrap your main code like this,

```go
package main

import (
  "github.com/mh-cbon/report-panic"
)

var MyRepo = "mh-cbon/demotest"

func main () {
  reportpanic.HandleMain(reportpanic.Gh(MyRepo), func () {
    panic("oh no!")
  })
}
```

## Install

```sh
glide get github.com/mh-cbon/report-panic
go get github.com/mh-cbon/report-panic
```

### Change the token reporter

To set a different token to report panics,
add this flag to your build compilation
`--ldflags "-X reportpanic.GhReporterToken=xxxxx"`


```sh
go build --ldflags "-X reportpanic.GhReporterToken=xxxxx"
```

### Change the template to generate the body report

1. Access to the templater instance of the reporter

```go
package main

import (
  "github.com/mh-cbon/report-panic"
)

func main () {
  reporter := reportpanic.Gh("your_user/your_repo")
  reporter.Templater.Template = "template {{content}} here"
  reportpanic.HandleMain(reporter, func () {
    panic("oh no!")
  })
}
```

2. Or, use ldflags at build time

```go
package main

import (
  "github.com/mh-cbon/report-panic"
)

// register your template in the init, or main func
func init(){
  panicreport.PanicReportTemplates["your"] = "template {{content}} here"
}

func main () {
  reportpanic.HandleMain(reportpanic.Gh("your_user/your_repo"), func () {
    panic("oh no!")
  })
}
```

At build time, use `ldflags` to configure the template

```sh
go build --ldflags "-X reportpanic.GhTemplate=your"
```

# Credits

[Mitchell Hashimoto](https://github.com/mitchellh) for the base code provided in [panicwrap](https://github.com/mitchellh/panicwrap)
