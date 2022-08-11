##Dumb thing which periodically parse everything you've specified in configuration

####Build darwin
```
GOOS=darwin go build -ldflags="-s -w" -o cmd/peparse/main cmd/peparse/main.go
```

####Build linux
```
GOOS=linux go build -ldflags="-s -w" -o cmd/peparse/main cmd/peparse/main.go
```

###Configuration structure

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

###Run
```
./cmd/peparse/main -help
```
