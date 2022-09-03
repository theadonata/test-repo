package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		host, _ := os.Hostname()
		c.JSON(http.StatusOK, gin.H{" test hostname": host})
	})

	r.Run()
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    // A very simple health check.
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")

    // In the future we could report back on the status of our DB, or our cache 
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(w, `{"alive": true}`)
}
