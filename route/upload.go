package route

import (
	// http router
	"github.com/gin-gonic/gin"

	// builtin
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func UploadSendFile(c *gin.Context) {
	resp := RESPONSE{}

	// Source
	file, err := c.FormFile("file")
	filename := c.PostForm("filename")
	if len(c.PostForm("filename")) < 1 {
		filename = file.Filename
	}

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

	if err := c.SaveUploadedFile(file, Config.File.Storage+filename); err != nil {
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

	resp.Data.Filename = filename
	resp.Data.Message = "Upload ok."

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	},
	)
}

func UploadByUrl(c *gin.Context) {
	resp := RESPONSE{}
	url := c.PostForm("url")
	filename := c.PostForm("filename")

	tokens := strings.Split(url, "/")
	if len(c.PostForm("filename")) < 1 {
		filename = tokens[len(tokens)-1]
	}
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
	resp.Data.Url = url

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	},
	)
}
