package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//设置路由分组
func main() {
	router := gin.Default()
	g1 := router.Group("/g1")
	g1.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "/g1/login ok")
	})
	g1.GET("/submit", func(c *gin.Context) {
		c.String(http.StatusOK, "/g1/submit ok")
	})
	g2 := router.Group("/g2")
	g2.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "/g2/login  ok")
	})
	g2.POST("submit", func(c *gin.Context) {
		c.String(http.StatusOK, "/g2/submit ok")
	})
	router.Run(":8000")
}
