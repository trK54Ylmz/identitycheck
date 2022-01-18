### KYC/AML check

The package provides a level of know your customer service for banking and other financial services.

Note: They may have some missing records. The source file will be populated from time to time.

#### Install

To start using identity check, install Go 1.16 or above. This package need Go modules to install. Run the following command to retrieve the library.

```bash
$ go get github.com/trk54ylmz/identitycheck/v0.1
```

#### Usage

The package needs name for filtering out results.

```go
ic, err := identitycheck.NewIdentityCheck()

name := "Barack Obama"

results, err := ic.Check(name)
```

Result objects are JSON serializable, example output will be like,

```json
[
    {
        "hash": "BRKBM",
        "name": "Barack Obama",
        "type": "pep",
        "birth_day": 4,
        "birth_month": 8,
        "birth_year": 1961,
        "alias": [
            "باراك أوباما",
            "Барак Обама",
            "Барак Абама",
            "贝拉克·奥巴马",
            "Baracus Obama",
            "ബറാക്ക് ഒബാമ",
            "Барақ Обама",
            "ባራክ ኦባማ"
        ],
        "birth_date": "1961-08-04T00:00:00Z",
        "country": {
            "code": "US",
            "iso": "USA",
            "name": "United States of America"
        }
    }
]
```
