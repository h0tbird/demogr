# demogr

[![GoReportCard Widget]][GoReportCard] [![Travis Widget]][Travis]

[GoReportCard]: https://goreportcard.com/report/h0tbird/demogr
[GoReportCard Widget]: https://goreportcard.com/badge/h0tbird/demogr
[Travis]: https://travis-ci.org/h0tbird/demogr
[Travis Widget]: https://travis-ci.org/h0tbird/demogr.svg?branch=master

Retrieves demographic data for a specified set of U.S. states from a public API and outputs that data in the requested format.

##### Install

```
go get -u github.com/h0tbird/demogr
```

##### Shell completion

```
eval "$(demogr --completion-script-${0#-})"
```

##### Help

```
demogr --help
usage: demogr [<flags>] <states>

Retrieves demographic data for a specified set of U.S. states from a public API and outputs that data in the requested format.

Flags:
  --help           Show context-sensitive help (also try --help-long and --help-man).
  --format="CSV"   Output-format parameter [CSV|averages]
  --max-workers=5  Maximum number of concurrent workers.

Args:
  <states>  Comma delimited list of U.S. states.
```

##### Usage examples

```
demogr --format CSV "Oklahoma,Indiana,New York,Rhode Island"
```
