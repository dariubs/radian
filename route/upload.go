package route

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// TODO: save file by specific filename
func UploadSendFile(c *gin.Context) {
	resp := RESPONSE{}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error: %s", err)

		resp.Error.HasError = true
		resp.Error.ErrorNumber = 1
		resp.Error.ErrorMessage = "Send incorrect file."
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"data": resp,
			},
		)
		return
	}

	if err := c.SaveUploadedFile(file, Config.File.Storage+file.Filename); err != nil {
		log.Printf("Error: %s", err)

		resp.Error.HasError = true
		resp.Error.ErrorNumber = 1
		resp.Error.ErrorMessage = "Could not be save."
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"data": resp,
			},
		)
		return
	}

	resp.Data.Filename = file.Filename
	resp.Data.Message = "Upload ok."

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	},
	)
}

func UploadByUrl(c *gin.Context) {
	resp := RESPONSE{}
	url := c.PostForm("url")

	tokens := strings.Split(url, "/")
	filename := tokens[len(tokens)-1]

	out, err := os.Create(Config.File.Storage + filename)
	if err != nil {
		resp.Error.HasError = true
		resp.Error.ErrorNumber = 1
		resp.Error.ErrorMessage = "File could not be saved."
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"data": resp,
			},
		)
		return
	}
	defer out.Close()

	response, err := http.Get(url)
	if err != nil {
		resp.Error.HasError = true
		resp.Error.ErrorNumber = 1
		resp.Error.ErrorMessage = "Can not get url content."
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"data": resp,
			},
		)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		resp.Error.HasError = true
		resp.Error.ErrorNumber = 1
		resp.Error.ErrorMessage = "Cannot write file."
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"data": resp,
			},
		)
		return
	}

	resp.Data.Filename = filename
	resp.Data.Message = "Upload ok."
	resp.Data.Filename = url

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	},
	)
}
