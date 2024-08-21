package api

import (
	"anime-manager/utils"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type RenameRequest struct {
	OldPath string `json:"oldPath"`
	NewPath string `json:"newPath"`
}

func magicHandler(c *gin.Context) {
	var req RenameRequest
	dir := os.Getenv("ROOT") + c.Param("path")
	err := clean(dir)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 绑定JSON请求体到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func clean(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("error reading directory %v: %v\n", dir, err)
		return err
	}

	for _, entry := range entries {
		i, _ := entry.Info()
		if i.IsDir() {
			continue
		}
		if wanted(i.Name()) {
			p, err := utils.Parse(i.Name())
			if err == nil {
				os.Rename(dir+"/"+i.Name(), fmt.Sprintf(`%s/%s S%02dE%02d%s%s`, dir, p.Title, p.Season, p.Episode, p.Language, p.Extension))
			}
			continue
		}
		err = os.Remove(dir + "/" + i.Name())
		if err != nil {
			return err
		}
	}

	return nil

}

func wanted(p string) bool {

	want := []string{".mp4", ".mkv", ".ssa", ".srt", ".ass"}

	ext := path.Ext(p)
	for _, v := range want {
		if ext == v {
			return true
		}
	}
	return false
}
