# demogr

[![License Widget]][License] [![GoReportCard Widget]][GoReportCard] [![Travis Widget]][Travis]

[License]: http://www.apache.org/licenses/LICENSE-2.0.txt
[License Widget]: https://img.shields.io/badge/license-APACHE2-1eb0fc.svg
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

##### Sample data:

```
curl -s "https://www.broadbandmap.gov/broadbandmap/census/state/Oklahoma?format=json" | jq .
```

```json
{
  "status": "OK",
  "responseTime": 31,
  "message": [],
  "Results": {
    "state": [
      {
        "geographyType": "STATE2010",
        "name": "Oklahoma",
        "fips": "40",
        "stateCode": "OK"
      }
    ]
  }
}
```

```
curl -s "https://www.broadbandmap.gov/broadbandmap/demographic/jun2014/state/ids/40?format=json" | jq .
```

```json
{
  "status": "OK",
  "responseTime": 89,
  "message": [],
  "Results": [
    {
      "geographyId": "40",
      "geographyName": "Oklahoma",
      "landArea": 66071.25626824,
      "population": 3884459,
      "households": 1755452,
      "raceWhite": 0.7766,
      "raceBlack": 0.0661,
      "raceHispanic": 0.0843,
      "raceAsian": 0.0122,
      "raceNativeAmerican": 0.0608,
      "incomeBelowPoverty": 0.169,
      "medianIncome": 47659.2346,
      "incomeLessThan25": 0.2887,
      "incomeBetween25to50": 0.2802,
      "incomeBetween50to100": 0.2956,
      "incomeBetween100to200": 0.1131,
      "incomeGreater200": 0.0223,
      "educationHighSchoolGraduate": 0.8175,
      "educationBachelorOrGreater": 0.2165,
      "ageUnder5": 0.0612,
      "ageBetween5to19": 0.2265,
      "ageBetween20to34": 0.2008,
      "ageBetween35to59": 0.3089,
      "ageGreaterThan60": 0.2026,
      "myAreaIndicator": false
    }
  ]
}
```
