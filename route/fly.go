package route

import (
	// image processor
	"github.com/disintegration/imaging"

	// http router
	"github.com/gin-gonic/gin"

	//  builtin
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ResizeThumbnailOnTheFly(c *gin.Context) {
	url := c.Query("url")

	urlresp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer urlresp.Body.Close()

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

	img, err := imaging.Decode(urlresp.Body)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

	imaging.Encode(c.Writer, img, 1)
}
