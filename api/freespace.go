package api

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
)

func freeSpaceHandler(c *gin.Context) {
	freeSpace, err := getDiskFreeSpace(os.Getenv("ROOT"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"free_space": freeSpace})
}

func getDiskFreeSpace(path string) (uint64, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, err
	}
	// 计算可用空间大小（字节）
	return stat.Bavail * uint64(stat.Bsize), nil
}
