package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/somejson", func(c *gin.Context) {
		names := []string{"zhang", "hello", "world"}
		c.SecureJSON(http.StatusOK, names)
	})
	router.Run(":8000")
}
