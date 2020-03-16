
# gotd - TD Ameritrade API in Go
[![GoDoc](https://godoc.org/github.com/z-Wind/gotd?status.png)](http://godoc.org/github.com/z-Wind/gotd)

## Table of Contents

* [Installation](#installation)
* [Examples](#examples)
* [Reference](#reference)

## Installation

    $ go get github.com/z-Wind/gotd

## Examples

### HTTP
```go
auth := NewAuth(redirectURL)
client := auth.GetClient(clientsecretPath)
td, err := New(client)
```

### HTTPS
```go
auth := NewAuth(redirectURL)
auth.SetTLS(TLSCertPath, TLSKeyPath)
client := auth.GetClient(clientsecretPath)
td, err := New(client)
```

### Quote
```go
service := NewQuotesService(td)
call := service.GetQuote("VTI")
quote, err := call.Do()
```

## Reference
- [TD Ameritrade API](https://developer.tdameritrade.com/apis)