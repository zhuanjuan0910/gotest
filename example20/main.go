package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {

			log.Println(file.Filename)
			// 上传文件到指定的路径
			c.SaveUploadedFile(file, "./copy_"+file.Filename)
			//c.SaveUploadedFile()
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8000")
}

																											
