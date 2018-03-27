package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB
	// router.Static("/", "./public")

	router.Use(DummyMiddleware())

	router.GET("/health", HandlerHealth)

	router.POST("/upload", HandlerUploadFile)
	router.POST("/file/create", HandlerCreateFile)
	router.GET("/file/:uuid", HandlerFileInfo)

	router.Run(":3000")
}

// DummyMiddleware DummyMiddleware
func DummyMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	fmt.Println("Init!")
	return func(c *gin.Context) {
		c.Next()
	}
}
