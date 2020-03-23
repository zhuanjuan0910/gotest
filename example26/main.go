package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type myForm struct{
	Colors []string `form:"colors[]"`
}
func test(c *gin.Context){
	var  postForm myForm
	if c.ShouldBind(&postForm)==nil{
		c.JSON(http.StatusOK,gin.H{"colors":postForm.Colors})
	}
}
	
func main(){
	router:=gin.Default()
	router.POST("/test",test)
	router.StaticFile("/testing","./abc.html")
	router.Run(":8000")
}