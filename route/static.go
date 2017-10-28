package route

import (
	// http router
	"github.com/gin-gonic/gin"

	// builtin
	"net/http"
)

// static pages
func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/upload")
}

func StaticUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"config": Config,
	})
}
