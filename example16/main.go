package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.POST("/form_post", func(c *gin.Context) {
        message := c.PostForm("message")
        nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

        c.JSON(http.StatusOK, gin.H{
            "status":  "posted",
            "message": message,
            "nick":    nick,
        })
    })
	router.Run(":8000")
	
}