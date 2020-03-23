package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user"  binding:"required"` //注意导出需要大写
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		var login LoginForm
		if c.ShouldBind(&login) == nil {
			if login.User == "tao" && login.Password == "bao" {
				c.String(http.StatusOK, "you are login")
			} else {
				c.String(401, "unauthorized")
			}
		}
	})
	router.Run(":8000")
}
