[![Go Report Card](https://goreportcard.com/badge/github.com/princjef/mageutil)](https://goreportcard.com/report/github.com/princjef/mageutil)
[![GitHub Actions](https://github.com/princjef/mageutil/workflows/Test/badge.svg)](https://github.com/princjef/mageutil/actions?query=workflow%3ATest+branch%3Amaster)
[![Release](https://img.shields.io/github/release/princjef/mageutil.svg)](https://github.com/princjef/mageutil/releases/latest)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/princjef/mageutil)

# mageutil

Set of utilities to make working with the [mage][] build system more seamless.
Each utility is in a subpackage of this one. You can install all of the
utilities by running:

```
go get github.com/princjef/mageutil
```

## Utilities

### bintool

```
github.com/princjef/mageutil/bintool
```

The `bintool` utility provides a way to manage binary tool dependencies in your
project automatically, ensuring that they are always the correct version and
keeping them namespaced within your project.

> NOTE: This utility is designed to optimize for tools that are built using
> [goreleaser][], but it should work well with any tool that can be downloaded
> in binary form from the internet. If you can't use this with one of your
> binary tools, please [open an issue][issues].

For each binary tool you have, you can instantiate a manager for it with
`bintool.New()`. For added convenience, you can store it as a variable using
`bintool.Must()` to panic on templating errors.

For example, the following setup will allow you to work with
[`golangci-lint`][golangci-lint] version 1.23.6:

```go
package main

import "github.com/princjef/mageutil/bintool"

var linter = bintool.Must(bintool.New(
	"golangci-lint{{.BinExt}}",
	"1.23.6",
	"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
))
```

Once this has been instantiated, you can perform operations like the following:

```go
// Check if the tool is installed for this project with the proper version
installed := linter.IsInstalled()

// Install the tool for this project
err := linter.Install()

// Install the tool for this project if it is not installed already or is the
// incorrect version
err := linter.Ensure()

// Execute the "run" command for the installed version of the linter.
err := linter.Command("run").Run()
```

For more detailed information, check out the [`bintool`][bintool] package docs.

### shellcmd

```
github.com/princjef/mageutil/shellcmd
```

The `shellcmd` utility provides a simple wrapper around a command to execute,
piping the output to the terminal and providing nicely-formatted logging.

Commands are built around the type `shellcmd.Command`, which represents a
command string as it would be typed into a terminal for natural typing and
reading. If you want a command to run tests on my package, you can do:

```go
package main

import "github.com/princjef/mageutil/shellcmd"

func Test() error {
	return shellcmd.Command("go test ./...").Run()
}
```

You can also run a list of commands in sequence easily using
`shellcmd.RunAll()`. If a command in the sequence fails, the commands that
follow it will not be run and the error will be returned. Since string literals
will be automatically interpreted as the appropriate type when passed to a
function, the strings can be passed directly in most cases

For example, to run your tests and then open coverage information in your web
browser:

```go
err := shellcmd.RunAll(
	"go test -coverprofile=coverage.out ./...",
	"go tool cover -html=coverage.out",
)
```

This can even be combined with a tool from the `bintool` package. Let's say you
want to run the linter from the example above prior to testing:

```go
err := shellcmd.RunAll(
	linter.Command("run"),
	"go test -coverprofile=coverage.out ./...",
	"go tool cover -html=coverage.out",
)
```

For more detailed information, check out the [`shellcmd`][shellcmd] package
docs.

[mage]: https://magefile.org/
[goreleaser]: https://goreleaser.com/
[golangci-lint]: https://github.com/golangci/golangci-lint
[bintool]: ./bintool/
[shellcmd]: ./shellcmd/
[issues]: https://github.com/princjef/mageutil/issues
