package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		name := c.DefaultQuery("name", "zhangsan")
		age := c.PostForm("age")
		message := c.DefaultPostForm("message", "hello")
		fmt.Printf("id=%s,name=%s,age=%s,message=%s", id, name, age, message)
	})
	router.Run(":8000")
}
