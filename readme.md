impossible [![Build Status](https://secure.travis-ci.org/daaku/impossible.png)](http://travis-ci.org/daaku/impossible) [![GoDoc](https://godoc.org/github.com/daaku/impossible?status.svg)](https://godoc.org/github.com/daaku/impossible) [![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](license)
==========

    import "github.com/daaku/impossible"

Documentation: https://godoc.org/github.com/daaku/impossible

Package impossible converts errors into panics. It is meant to be used when the
error is impossible.

## Usage

#### func  Error

```go
func Error(err error)
```
Error panics with the error if it is not nil.
