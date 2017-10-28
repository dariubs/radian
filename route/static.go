package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// static page
func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/upload")
}

func StaticUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"config": Config,
	})
}
