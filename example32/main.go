package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cookies", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookies")
		if err != nil {
			cookie = "notset"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("cookies  is %s", cookie)
	})
	router.Run(":8000")
}
