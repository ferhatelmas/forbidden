# Forbidden Usernames in Golang

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/forbidden)
[![Build Status](https://travis-ci.org/ferhatelmas/forbidden.svg?branch=master)](https://travis-ci.org/ferhatelmas/forbidden)

> A handy list of usernames to forbid in Golang.

## Install

By go tool: `go get github.com/ferhatelmas/forbidden`

## Usage

```go
import github.com/ferhatelmas/forbidden

forbidden.Name("Admin") // true
forbidden.Name(".json") // true
```

## RELATED

Actual list is collected at [dsignr/disallowed-usernames](https://github.com/dsignr/disallowed-usernames).

## CONTRIBUTING

Feel free to create a PR with additions or removal requests [here](https://github.com/ferhatelmas/forbidden/issues/new) or above related repo. For sure, don't forget adding explanation of request.

## LICENSE

MIT © [Ferhat Elmas](https://github.com/ferhatelmas)
