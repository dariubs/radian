package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func DeleteFile(c *gin.Context) {
	fileName := c.Param("filename")
	var err = os.Remove(Config.File.Storage + fileName)
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
	err := os.Rename(Config.File.Storage+fileName, Config.File.Storage+newName)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	fmt.Println("==> done renaming file")
	c.String(http.StatusOK, fmt.Sprintf("File %s renamed to %s.", fileName, newName))
}
