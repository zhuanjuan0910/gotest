package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./b.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "b.html", gin.H{ //注意此处写文件名
			"title": "bnskjcnnck",
		})
	})
	router.Run(":8000")
}
