[![Go Report Card](https://goreportcard.com/badge/github.com/jbowles/disfun)](https://goreportcard.com/report/github.com/jbowles/disfun)
[![GoDoc](https://godoc.org/github.com/jbowles/disfun?status.svg)](https://godoc.org/github.com/jbowles/disfun)

# DisFun
Inspired by the Ruby Gem [Distance-Measures](https://github.com/reddavis/Distance-Measures) and [Measurable](https://github.com/agarie/measurable) and my need to have available a number of distance functions for things like k-Nearest-Neighbor, k-Means, string metrics, and the like....

## Algos
I'll add algorithm implementations as they arise. Some imlementations have been forked from existing projects where I have made modifcations to suite my purposes (see external_licenses).

# Test Coverage
coverage: 87.3% of statements

## Run benchmarks

```sh
go test -bench .

## or

go test -bench=.

## or for levenshtien

go test -bench Lev

## or for euclidean
go test -bench Euc

## or use gobenchui
gobenchui .
```
