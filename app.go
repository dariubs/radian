package main

import (
	// config
	"github.com/BurntSushi/toml"
	"github.com/dariubs/radian/config"

	// http router
	"github.com/gin-gonic/gin"

	// builtin
	"fmt"
	"log"
)

var err error
var router *gin.Engine
var Config config.CONFIG

func main() {
	// load config file
	if _, err = toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
		return
	}

	// router
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// static contents
	router.Static("/public", "./public")

	// go templates
	router.LoadHTMLGlob("view/*")

	// routes
	// TODO: add authentication
	router.GET("/", Index)

	upload := router.Group("/upload")
	{
		upload.POST("/sendfile", UploadSendFile)
		upload.POST("/byurl", UploadByUrl)
	}

	router.Static("/show", Config.File.Storage)

	modify := router.Group("/modify")
	{
		modify.DELETE("/delete/:filename", DeleteFile)
		modify.PATCH("/patch/:filename", RenameFile)
	}

	resize := router.Group("/resize")
	{
		resize.GET("/thumbnail/:width/:height/:filename", ResizeThumbnail)
		resize.GET("/fit/:width/:height/:filename", ResizeFit)
		resize.GET("/fill/:width/:height/:filename", ResizeFill)
	}

	// run app
	router.Run(fmt.Sprintf(Config.Server.Port))
}
