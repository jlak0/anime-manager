package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func getDirHandler(c *gin.Context) {

	v, _ := getDir(c.Param("path"))
	c.JSON(http.StatusOK, v)
}

func getDir(dir string) ([]File, error) {
	dir = os.Getenv("ROOT") + dir

	f := []File{}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("error reading directory %v: %v\n", dir, err)
		return f, err
	}

	for _, entry := range entries {
		i, _ := entry.Info()
		size := i.Size()
		if i.IsDir() {
			size, err = getDirSize(dir + "/" + i.Name())
			if err != nil {
				size = 0
			}
		}
		f = append(f, File{
			Name:  i.Name(),
			Path:  dir,
			Size:  size,
			IsDir: i.IsDir(),
		})
	}
	return f, nil
}

func getDirSize(dirPath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 累加每个文件的大小
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}
