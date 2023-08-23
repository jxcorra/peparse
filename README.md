Dumb thing which periodically parse everything you've specified in configuration
![Coverage](https://img.shields.io/badge/Coverage-80.0%25-brightgreen)

![lint](https://github.com/jxcorra/peparse/actions/workflows/lint.yaml/badge.svg)
![build](https://github.com/jxcorra/peparse/actions/workflows/build.yaml/badge.svg)
![tests](https://github.com/jxcorra/peparse/actions/workflows/test.yaml/badge.svg)
![cov](https://github.com/jxcorra/peparse/wiki/coverage.svg)

Build darwin
```
GOOS=darwin go build -ldflags="-s -w" -o cmd/peparse/main cmd/peparse/main.go
```

Build linux
```
GOOS=linux go build -ldflags="-s -w" -o cmd/peparse/main cmd/peparse/main.go
```

Configuration structure

- `resources` array - resource objects for parallel parse
```
# Example
{
    "resources": [
        {
            "url": "<resource url>",
            "search": [
                {
                    "key": {
                        "element": "a",
                        "class": "<class>"
                    },
                    "withText": true,  # when text of token is needed
                    "parse": [  # array of attributes which values need to grab
                        {
                            "attr": "href"
                        }
                    ]
                }
            ]
        },
        ...
    ]
}
```

Run
```
./cmd/peparse/main -help
```
