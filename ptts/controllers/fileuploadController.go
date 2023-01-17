package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// "/fileupload"
func FileuploadGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "fileupload.html", gin.H{
			"title":      "FileUpload",
			"fileupload": "active",
			"content":    "",
		})
	}
}
