package route

import (
	// image processor
	"github.com/disintegration/imaging"

	// http router
	"github.com/gin-gonic/gin"

	// "io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ShowOnTheFly(c *gin.Context) {
	url := c.Query("url")

	urlresp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: %s", err)

		img, _ := imaging.Open(Config.File.Default)

		imaging.Encode(c.Writer, img, 1)
		return
	}
	defer urlresp.Body.Close()

	img, err := imaging.Decode(urlresp.Body)
	if err != nil {
		log.Fatalf("Open failed: %v", err)

		img, _ := imaging.Open(Config.File.Default)

		imaging.Encode(c.Writer, img, 1)
		return
	}

	imaging.Encode(c.Writer, img, 1)
}

func ResizeThumbnailOnTheFly(c *gin.Context) {
	url := c.Query("url")

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

	urlresp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: %s", err)

		img, _ := imaging.Open(Config.File.Default)

		img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

		imaging.Encode(c.Writer, img, 1)
		return
	}
	defer urlresp.Body.Close()

	img, err := imaging.Decode(urlresp.Body)
	if err != nil {
		log.Fatalf("Open failed: %v", err)

		img, _ := imaging.Open(Config.File.Default)

		img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

		imaging.Encode(c.Writer, img, 1)
		return
	}

	img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

	imaging.Encode(c.Writer, img, 1)
}
