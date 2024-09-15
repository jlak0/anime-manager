package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {

	// 设置文件上传的路径
	uploadPath :=  c.PostForm("uploadPath")
	if uploadPath == "" {
		fmt.Println("uploadPath is empty")
		return
	}
	uploadPath = os.Getenv("ROOT") + uploadPath


	// 获取多文件
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("err", err)
		return
	}

	// 获取文件列表
	files := form.File["files"]
	fmt.Println("files", len(files))
	// 遍历所有上传的文件
	for _, file := range files {
		// 拼接文件保存路径
		filename := filepath.Join(uploadPath, filepath.Base(file.Filename))

		// 保存文件到服务器
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println("err", err)
			return
		}
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d files uploaded!", len(files))})

}
