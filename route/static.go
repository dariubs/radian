package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// index page
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"config": Config,
	})
}
