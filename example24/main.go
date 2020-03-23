package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func test(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil { //绑定get参数或者post参数
		log.Println(person.Name)
		log.Println(person.Address)

	}
	c.String(http.StatusOK, "success")
}
func main() {
	router := gin.Default()
	router.GET("/test", test)
	router.POST("/testing", test)
	router.Run(":8000")
}
