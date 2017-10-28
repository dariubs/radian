package route

import (
	//  http router
	"github.com/gin-gonic/gin"

	// builtin
	"log"
	"net/http"
	"os"
)

func DeleteFile(c *gin.Context) {
	resp := RESPONSE{}

	filename := c.Param("filename")

	var err = os.Remove(Config.File.Storage + filename)
	if err != nil {
		log.Printf("Error: %s", err)

		resp.Error.HasError = true
		resp.Error.ErrorNumber = 1
		resp.Error.ErrorMessage = "Cannot remove file."
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"data": resp,
			},
		)
		return
	}

	resp.Data.Filename = filename
	resp.Data.Message = "Remove ok."

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func RenameFile(c *gin.Context) {
	resp := RESPONSE{}

	filename := c.Param("filename")
	newname := c.Param("newname")

	var err = os.Rename(Config.File.Storage+filename, Config.File.Storage+newname)
	if err != nil {
		log.Printf("Error: %s", err)

		resp.Error.HasError = true
		resp.Error.ErrorNumber = 1
		resp.Error.ErrorMessage = "Cannot rename file."
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"data": resp,
			},
		)
		return
	}

	resp.Data.Filename = newname
	resp.Data.Message = "Rename ok."

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
