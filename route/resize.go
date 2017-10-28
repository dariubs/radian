package route

import (
	// image processor
	"github.com/disintegration/imaging"

	// http router
	"github.com/gin-gonic/gin"

	//  builtin
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func IsFileExist(file string) bool {
	if _, err := os.Stat(Config.File.Storage + file); os.IsNotExist(err) {
		return false
	}
	return true
}

func ResizeThumbnail(c *gin.Context) {
	filename := c.Param("filename")

	width, err := strconv.Atoi(c.Param("width"))
	if err != nil {
		log.Printf("Error: %s", err)
		width = 0
	}

	height, err := strconv.Atoi(c.Param("height"))
	if err != nil {
		log.Printf("Error: %s", err)
		height = 0
	}

	exist := IsFileExist(filename)

	if exist {
		img, err := imaging.Open(Config.File.Default)

		img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

		imaging.Encode(c.Writer, img, 1)
		return
	}

	img, err := imaging.Open(Config.File.Storage + filename)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

	imaging.Encode(c.Writer, img, 1)
}

func ResizeFit(c *gin.Context) {
	filename := c.Param("filename")

	width, err := strconv.Atoi(c.Param("width"))
	if err != nil {
		log.Printf("Error: %s", err)
		width = 0
	}

	height, err := strconv.Atoi(c.Param("height"))
	if err != nil {
		log.Printf("Error: %s", err)
		height = 0
	}

	exist := IsFileExist(filename)

	if exist {
		img, err := imaging.Open(Config.File.Default)

		img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

		imaging.Encode(c.Writer, img, 1)
		return
	}

	img, err := imaging.Open(Config.File.Storage + filename)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	img = imaging.Fit(img, width, height, imaging.Lanczos)

	imaging.Encode(c.Writer, img, 1)
}

func ResizeFill(c *gin.Context) {
	filename := c.Param("filename")

	width, err := strconv.Atoi(c.Param("width"))
	if err != nil {
		log.Printf("Error: %s", err)
		width = 0
	}

	height, err := strconv.Atoi(c.Param("height"))
	if err != nil {
		log.Printf("Error: %s", err)
		height = 0
	}

	exist := IsFileExist(filename)

	if exist {
		img, err := imaging.Open(Config.File.Default)

		img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

		imaging.Encode(c.Writer, img, 1)
		return
	}

	img, err := imaging.Open(Config.File.Storage + filename)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	img = imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)

	imaging.Encode(c.Writer, img, 1)
}
