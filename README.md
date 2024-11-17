# contextx

[![License](https://img.shields.io/github/license/rakunlabs/contextx?color=red&style=flat-square)](https://raw.githubusercontent.com/rakunlabs/contextx/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/rakunlabs_contextx?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=rakunlabs_contextx)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/rakunlabs/contextx/test.yml?branch=main&logo=github&style=flat-square&label=ci)](https://github.com/rakunlabs/contextx/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/rakunlabs/contextx?style=flat-square)](https://goreportcard.com/report/github.com/rakunlabs/contextx)
[![Go PKG](https://raw.githubusercontent.com/rakunlabs/.github/main/assets/badges/gopkg.svg)](https://pkg.go.dev/github.com/rakunlabs/contextx)

Context value helper library.  
Differences between context package this holds value in a map with mutex lock. Values readable from parent context.

```sh
go get github.com/rakunlabs/contextx
```

## Usage

```go
// set value first initialize context value map if not exist
ctx := contextx.WithValue(context.Background(), "secret", "xxx")

// map like access
if v, ok := contextx.Value[string](ctx, "secret"); ok {
    // v is string
}
```
