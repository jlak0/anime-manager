package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type FileToDelete struct {
	Name string
}

func deleteHandler(c *gin.Context) {
	var list []FileToDelete
	dir := c.Param("path")

	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.Split(strings.Trim(dir, "/"), "/")) <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不能删除根目录下的文件夹", "path": dir})
		return
	}
	for _, f := range list {
		err := os.RemoveAll(fmt.Sprintf("%s%s/%s", os.Getenv("ROOT"), dir, f.Name))
		fmt.Printf("%s%s/%s", os.Getenv("ROOT"), dir, f.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
