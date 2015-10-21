#DisFun
Inspired by the Ruby Gem [Distance-Measures](https://github.com/reddavis/Distance-Measures) and [Measurable[(https://github.com/agarie/measurable) and my need to have available a number of distance functions for things like k-Nearest-Neighbor, k-Means, string metrics, and the like....

## Run tests

## Run benchmarks
Desciription of testing flags: https://golang.org/cmd/go/#hdr-Description_of_testing_flags

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
