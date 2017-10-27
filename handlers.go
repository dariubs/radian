package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/disintegration/imaging"
	"io"
	"net/http"
	"os"
	"strings"
	"strconv"
	"log"
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

	if err := c.SaveUploadedFile(file, Config.File.Storage + file.Filename); err != nil {
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
	err := os.Rename(fileName, newName)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	fmt.Println("==> done renaming file")
	c.String(http.StatusOK, fmt.Sprintf("File %s renamed to %s.", fileName, newName))
}

func ResizeThumbnail(c *gin.Context)  {
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

	img, err := imaging.Open(Config.File.Storage + filename)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	img = imaging.Thumbnail(img, width, height, imaging.CatmullRom)

	imaging.Encode(c.Writer, img, 1)
}

func ResizeFit(c *gin.Context)  {
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

	img, err := imaging.Open(Config.File.Storage + filename)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	img = imaging.Fit(img, width, height, imaging.Lanczos)

	imaging.Encode(c.Writer, img, 1)
}

func ResizeFill(c *gin.Context)  {
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

	img, err := imaging.Open(Config.File.Storage + filename)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	img = imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)

	imaging.Encode(c.Writer, img, 1)
}
