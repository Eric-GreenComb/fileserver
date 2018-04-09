package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	if ServerConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB
	// router.Static("/", "./public")

	router.Use(favicon.New("./fav.ico"))

	router.LoadHTMLGlob("templates/*")

	router.Use(DummyMiddleware())

	router.GET("/", HandlerIndex)
	router.GET("/health", HandlerHealth)

	router.GET("/upload", HandlerGetUploadFile)
	router.POST("/upload", HandlerUploadFile)
	router.POST("/file/create", HandlerCreateFile)
	router.GET("/file/:uuid", HandlerFileInfo)

	router.GET("/test", HandlerGetTest)
	router.POST("/test", HandlerTest)

	router.Run(":" + ServerConfig.Port)
}

// DummyMiddleware DummyMiddleware
func DummyMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	fmt.Println("Init!")
	return func(c *gin.Context) {
		c.Next()
	}
}
