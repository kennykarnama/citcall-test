## citcall-test

This simple app is for citcall test for transforming json data to HTML

This app will serve http 

Endpoints

```
/countries

Response: html page

Error Response:

204, No Content if json response is empty
500, Internal server error if other error is encountered
```

## Usage

```
Usage: citcall-test.exe [--port PORT] [--external-request-timeout EXTERNAL-REQUEST-TIMEOUT] [--citcall-base-url CITCALL-BASE-URL]

Options:
  --port PORT, -p PORT   port for serving http server [default: 8080]
  --external-request-timeout EXTERNAL-REQUEST-TIMEOUT
                         http request timeout for external client [default: 15s]
  --citcall-base-url CITCALL-BASE-URL
                         citcall base url [default: https://citcall.com]
  --help, -h             display this help and exit
```

## Approach

- App calls endpoint `/test/countries.json`
- Parse as struct
- Render as view

## Unit test

To run unit test, please run `go test ./... -v`

## Dependency library

- github.com/alexflint/go-arg v1.4.3
- github.com/stretchr/testify v1.7.0
- https://github.com/vektra/mockery --> for mock generation

## Improvements

- Pagination if JSON response is large
- Streaming write 
