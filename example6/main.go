package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusOK, "https://www.baidu.com")
	})
	router.Run(":8000")
}
