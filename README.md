
# gotd - TD Ameritrade API in Go
[![GoDoc](https://godoc.org/github.com/z-Wind/gotd?status.png)](http://godoc.org/github.com/z-Wind/gotd)

## Table of Contents

* [Apply](#apply)
* [Installation](#installation)
* [Examples](#examples)
* [Reference](#reference)

## Apply
- Go to [TD Ameritrade API](https://developer.tdameritrade.com/apis)
- Follow this site  [[TD] API 申請流程](https://zwindr.blogspot.com/2018/10/td-api.html)
- Modify ***client_secret.json.sample*** and pass it to auth.GetClient

## Installation

    $ go get github.com/z-Wind/gotd

## Examples

### HTTP
```go
auth := NewAuth()
client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")
td, err := New(client)
```

### HTTPS
```go
auth := NewAuth()
auth.SetTLS(TLSCertPath, TLSKeyPath)
client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")
td, err := New(client)
```

### Quotes
```go
call := td.Quotes.GetQuote("VTI")
quote, err := call.Do()
```

## Reference
- [TD Ameritrade API](https://developer.tdameritrade.com/apis)
