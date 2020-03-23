package main

import "github.com/gin-gonic/gin"

func PostParam(c *gin.Context) {
	message := c.PostForm("message")
	nickname := c.DefaultPostForm("nickname", "abc")
	c.JSON(200, gin.H{
		"message":  message,
		"nickname": nickname,
	})
}
func main() {
	router := gin.Default()
	router.POST("/appPost", PostParam)
	router.StaticFile("/file", "./abc.html")
	router.Run(":8000")
}
