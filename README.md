# WRequestID

![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wrequestid?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wrequestid)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wrequestid)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wrequestid/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wrequestid)](https://goreportcard.com/report/github.com/wyy-go/wrequestid)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wrequestid?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wrequestid/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wrequestid)


Request ID middleware for Gin Framework. Adds an indentifier to the response using the `X-Request-ID` header. Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.

## Example

```go
package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wrequestid"
	"log"
	"net/http"
)

func main() {
	r := gin.New()

	r.Use(wrequestid.New())

	r.GET("/", func(c *gin.Context) {
		var buf bytes.Buffer
		buf.WriteString("GetRequestID:\t" + wrequestid.GetRequestID(c))
		buf.WriteString("\n")
		buf.WriteString("FromRequestID:\t" + wrequestid.FromRequestID(c.Request.Context()))
		c.String(http.StatusOK, buf.String())
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

```
