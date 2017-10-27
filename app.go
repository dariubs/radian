package main

import (
	// config
	"github.com/BurntSushi/toml"
	"github.com/dariubs/radian/config"

	// http router
	"github.com/dariubs/radian/route"
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
	router.GET("/", route.Index)

	upload := router.Group("/upload")
	{
		upload.POST("/sendfile", route.UploadSendFile)
		upload.POST("/byurl", route.UploadByUrl)
	}

	router.Static("/show", Config.File.Storage)

	modify := router.Group("/modify")
	{
		modify.DELETE("/delete/:filename", route.DeleteFile)
		modify.PATCH("/patch/:filename", route.RenameFile)
	}

	resize := router.Group("/resize")
	{
		resize.GET("/thumbnail/:width/:height/:filename", route.ResizeThumbnail)
		resize.GET("/fit/:width/:height/:filename", route.ResizeFit)
		resize.GET("/fill/:width/:height/:filename", route.ResizeFill)
	}

	// run app
	router.Run(fmt.Sprintf(Config.Server.Port))
}
