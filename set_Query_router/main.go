package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取query string的参数，即url中疑问号后面的参数
func GetPara(c *gin.Context) {
	nickname := c.Query("nickname")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "nickname=%s,passwd=%s get query ok", nickname, passwd)
}
func PostPara(c *gin.Context) {
	nickname := c.Query("nickname")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "nickname=%s,passwd=%s post query  ok", nickname, passwd)
}

func main() {
	router := gin.Default()
	router.GET("/appGet", GetPara)
	router.POST("/appPost", PostPara)

	router.Run(":8000")

}
