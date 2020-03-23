package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPara(c *gin.Context) {
	nickname := c.Param("nickname")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "nickname=%s,passwd=%s get param ok", nickname, passwd)
}
func PostPara(c *gin.Context) {
	nickname := c.Param("nickname")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "nickname=%s,passwd=%s post param ok", nickname, passwd)
}

func main() {
	router := gin.Default()
	router.GET("/appGet/:nickname/:passwd", GetPara) //':'必须要匹配,'*'选择匹配,即存在就匹配,否则可以不考虑
	router.POST("/appPost/:nickname/*passwd", PostPara)

	router.Run(":8000")

}
