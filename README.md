API Auth for Gin (Go Framework)
====================================

A simple random token generator. This is a port of [TokenGenerator](https://github.com/wernerdweight/TokenGenerator/tree/master) for PHP.

[![Build Status](https://www.travis-ci.com/wernerdweight/token-generator-go.svg?branch=master)](https://www.travis-ci.com/wernerdweight/token-generator-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/wernerdweight/token-generator-go)](https://goreportcard.com/report/github.com/wernerdweight/token-generator-go)
[![GoDoc](https://godoc.org/github.com/wernerdweight/token-generator-go?status.svg)](https://godoc.org/github.com/wernerdweight/token-generator-go)
[![go.dev](https://img.shields.io/badge/go.dev-pkg-007d9c.svg?style=flat)](https://pkg.go.dev/github.com/wernerdweight/token-generator-go)


Installation
------------

### 1. Installation

```bash
go get github.com/wernerdweight/token-generator-go
```

Usage
------------

```go
import "github.com/wernerdweight/token-generator-go"

// default
g := NewTokenGenerator("") // "" - default alphabet
token := g.Generate(32)    // produces a 32-character token
// token would be something like "Ghll-NpLhRcWPHCGV7Ix-O2LFYSeKsES"

// custom alphabet
g := NewTokenGenerator("abc") // token will contain only characters from the alphabet "abc" 

// custom length
g := NewTokenGenerator("")
token := g.Generate(64) // produces a 64-character token
```

License
-------
This package is under the MIT license. See the complete license in the root directory of the bundle.
