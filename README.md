# contextx

[![License](https://img.shields.io/github/license/worldline-go/contextx?color=red&style=flat-square)](https://raw.githubusercontent.com/worldline-go/contextx/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/worldline-go_contextx?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=worldline-go_contextx)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/worldline-go/contextx/test.yml?branch=main&logo=github&style=flat-square&label=ci)](https://github.com/worldline-go/contextx/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/worldline-go/contextx?style=flat-square)](https://goreportcard.com/report/github.com/worldline-go/contextx)
[![Go PKG](https://raw.githubusercontent.com/worldline-go/guide/main/badge/custom/reference.svg)](https://pkg.go.dev/github.com/worldline-go/contextx)

Context value helper library.  
Differences between context package this holds value in a map with mutex lock.

```sh
go get github.com/worldline-go/utility/contextx
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
