package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileToRename struct {
	NewPath string `json:"newPath"`
	OldPath string `json:"oldPath"`
}

func renameHandler(c *gin.Context) {
	var data FileToRename
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.NewPath == data.OldPath {
		c.JSON(http.StatusNoContent, gin.H{"status": "No Change"})
		return
	}
	fmt.Println(os.Getenv("ROOT") + data.OldPath)
	err := os.Rename(os.Getenv("ROOT")+data.OldPath, os.Getenv("ROOT")+data.NewPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
