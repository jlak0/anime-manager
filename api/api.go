package api

import (
	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

func Serve() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	r.GET("/api/dir/*path", getDirHandler)
	r.DELETE("/api/dir/*path", deleteHandler)
	r.PATCH("/api/dir/*path", magicHandler)

	r.PATCH("/api/rename", renameHandler)
	r.GET("/api/freespace", freeSpaceHandler)
	r.POST("/api/upload", uploadHandler)
	r.Run(":8070")
}
