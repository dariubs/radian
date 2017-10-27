package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"config": Config,
	})
}

func UploadSendFile(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
}

func UploadByUrl(c *gin.Context) {
	url := c.PostForm("url")

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	output, err := os.Create(Config.File.Storage + fileName)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s.", url, fileName))

}

func DeleteFile(c *gin.Context) {
  fileName := c.Param("filename")
  var err = os.Remove(fileName)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	fmt.Println("==> done deleting file")
  c.String(http.StatusOK, fmt.Sprintf("File %s Deleted.", fileName))
}

func RenameFile(c *gin.Context) {
  fileName := c.Param("filename")
	newName := c.Param("newname")
  err :=  os.Rename(filename, newname)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	fmt.Println("==> done renaming file")
  c.String(http.StatusOK, fmt.Sprintf("File %s renamed to %s.", fileName, newName))
}
