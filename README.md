# go-sanitize
**go-sanitize** implements a simple library of sanitation methods for data sanitation and reduction. This package was inspired by processing and protecting incoming user generated content while ensuring the data will be in the correct format. This project follows Go best practices and you can view the standards and specifications at the [end of this readme](https://github.com/mrz1836/go-sanitize#adheres-to-effective-go-standards).

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-sanitize?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-sanitize)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/b11a08d5619849a0ae911d91e3bb47c7)](https://www.codacy.com/app/mrz1818/go-sanitize?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-sanitize&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-sanitize.svg?branch=master)](https://travis-ci.com/mrz1836/go-sanitize)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-sanitize.svg?style=flat)](https://github.com/mrz1836/go-sanitize/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-sanitize?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-sanitize) |


## Table of Contents
- [Installation](https://github.com/mrz1836/go-sanitize#installation)
- [Documentation](https://github.com/mrz1836/go-sanitize#documentation)
- [Examples & Tests](https://github.com/mrz1836/go-sanitize#examples--tests)
- [Benchmarks](https://github.com/mrz1836/go-sanitize#benchmarks)
- [Code Standards](https://github.com/mrz1836/go-sanitize#code-standards)
- [Usage](https://github.com/mrz1836/go-sanitize#usage)
- [Maintainers](https://github.com/mrz1836/go-sanitize#maintainers)
- [Contributing](https://github.com/mrz1836/go-sanitize#contributing)
- [License](https://github.com/mrz1836/go-sanitize#license)

## Installation

**go-sanitize** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-sanitize
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-sanitize).

## Examples & Tests
All unit tests and [examples](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-sanitize) and uses [Go version 1.11.x](https://golang.org/doc/go1.11). View the [deployment configuration file](https://github.com/mrz1836/go-sanitize/blob/master/.travis.yml).
```bash
$ cd ../go-sanitize
$ go test ./... -v
```

## Benchmarks
Run the Go [benchmarks](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go):
```bash
$ cd ../go-sanitize
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](https://github.com/mrz1836/go-sanitize/blob/master/CODE_STANDARDS.md).

## Usage
- View the [examples](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go)
- View the [benchmarks](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go)
- View the [tests](https://github.com/mrz1836/go-sanitize/blob/master/sanitize_test.go)

Basic implementation:
```golang
package main

import (
	"fmt"
	"github.com/mrz1836/go-sanitize"
)

func main() {

	//Execute and print
	fmt.Println(sanitize.IPAddress("  ##!192.168.0.1!##  "))

	// Output: 192.168.0.1
}
```

## Maintainers

[@MrZ1836](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](https://github.com/mrz1836/go-sanitize/blob/master/CONTRIBUTING.md) and follow the [code of conduct](https://github.com/mrz1836/go-sanitize/blob/master/CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin%20cash-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-sanitize)

## License

![License](https://img.shields.io/github/license/mrz1836/go-sanitize.svg?style=flat)
