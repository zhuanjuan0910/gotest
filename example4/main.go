package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 在单个路由中使用 MyBenchLogger 中间件
	r.GET("/benchmark", MyBenchLogger(), func(c *gin.Context) {
		c.String(http.StatusOK, "/benchmark ok")
		fmt.Println("here")
	})

	authorized := r.Group("/")
	// 在路由组中使用中间件
	authorized.Use(func(c *gin.Context) {})
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "/login OK")
		})
		authorized.POST("/submit", func(c *gin.Context) {})
		authorized.POST("/read", func(c *gin.Context) {})

		// testing group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			c.String(http.StatusOK, "/testing/analytics OK")
		})
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8000")
}

// 中间件示例
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next() //先调下一个handler执行后再执行下面的代码
		fmt.Println("after middleware")
	}
}

//控制台打印结果为
// before middleware
// here
// after middleware
