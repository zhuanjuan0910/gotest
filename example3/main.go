package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//静态文件服务
func main() {
	router := gin.Default()
	router.Static("/file", "../example1")               //目录映射
	router.StaticFile("/file1", "../example1/abc.html") //文件映射
	router.StaticFS("/showDir", http.Dir("."))
	router.Run(":8000")
}
