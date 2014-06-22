go-commons-validator
====================

Go port of the Apache Commons Validator framework

A work in progress

## Installation

Simply install the package to your [$GOPATH](http://code.google.com/p/go-wiki/wiki/GOPATH "GOPATH") with the [go tool](http://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/dsparling/go-commons-validator
```
Make sure [Git is installed](http://git-scm.com/downloads) on your machine and in your system's `PATH`.

*`go get` installs the latest tagged release*

## Examples

[Example.go](https://github.com/dsparling/go-commons-validator/blob/master/examples/example.go) is a sort of hello world for go-commons-validator and should get you started for the barebones necessities of using the package.

	cd examples
	go run example.go

## Email

	// true
	fmt.Println(email.Validate("test@example.com"))

	// false
	fmt.Println(email.Validate("testexample.com"))
