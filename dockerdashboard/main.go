package main

import (
	"dockerdashboard/example"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./web")

	r.GET("/api/v1/container/list", func(ctx *gin.Context) {
		containerInfos, err := example.List()
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err,
			})
		} else {
			ctx.JSON(200, containerInfos)
		}

	})
	r.Run(":30080")
}
