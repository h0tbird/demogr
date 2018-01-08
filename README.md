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
  --help                  Show context-sensitive help (also try --help-long and --help-man).
  --format="CSV"          Output-format parameter [CSV|averages]
  --max-workers=5         Maximum number of concurrent workers.
  --data-version=jun2014  Specify the data version.

Args:
  <states>  Comma delimited list of U.S. states.
```

##### Usage examples

```
demogr --format CSV "Oklahoma,Indiana,New York,Rhode Island"
Indiana,6640448,2915239,0.154200,51519.390000
New York,19569355,8337971,0.152400,65123.585700
Oklahoma,3884459,1755452,0.169000,47659.234600
Rhode Island,1050223,471217,0.140000,60530.681800
```

```
demogr --data-version dec2013 --format CSV --max-workers 2 "Oklahoma,Indiana,New York,Rhode Island"
Indiana,6606176,2884150,0.154400,51410.505000
New York,19531372,8279920,0.152400,64766.918100
Oklahoma,3855947,1732748,0.169200,47429.411000
Rhode Island,1051689,468983,0.140100,60172.335800
```

##### Asciinema demo

[![asciicast](https://asciinema.org/a/oGyYiDYxmKwYoUDjwkTqhqL4A.png)](https://asciinema.org/a/oGyYiDYxmKwYoUDjwkTqhqL4A)

##### Assumptions

* Requests return a small amount of data and pagination is not needed.
* Fail fast: API error retries, exponential backoff and cirquit breakers are not implemented.
* Defaults to `--format CSV --max-workers 5 --data-version jun2014`.
* I am weighting the `incomeBelowPoverty` by the `population` in each state.

##### Census data sample

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

##### Demographic data sample

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
