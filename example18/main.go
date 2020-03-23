package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "./copy_h.txt")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})
	router.Run(":8000")
}
