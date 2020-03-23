package main

import(
	 "github.com/gin-gonic/gin"
	 "net/http"
)

func main(){
	router:=gin.Default()
	router.GET("/appGet",func(c *gin.Context){
		c.String(http.StatusOK,"hello world")
	})
	router.Run(":8000")
}