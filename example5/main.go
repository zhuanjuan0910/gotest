package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./abc.html")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "abc.html", gin.H{
			"title": "GIN:测试加载模板",
		})
	})
	router.Run(":8000")
}
