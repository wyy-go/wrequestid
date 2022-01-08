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
