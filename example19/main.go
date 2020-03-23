package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个不包含中间件的路由器
	r := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	r.Use(gin.Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 路由添加中间件，可以添加任意多个
	r.GET("/benchmark", MyBenchLogger(), func(c *gin.Context) {
		fmt.Println("hello world")
		c.String(http.StatusOK, "bchdbdhbkh")
	})

	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "you are login")
		})
		authorized.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "you are submit")
		})
		authorized.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "you are read")
		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			c.String(http.StatusOK, "I am analytics")
		})
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Next()
		fmt.Println("after middleware")
	}
}
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("how are you")
	}
}
