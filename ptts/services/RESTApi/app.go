package servicies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello Get %s", name)
	}
}

func ApiPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello Post %s", name)
	}
}

func ApiPutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello Put %s", name)
	}
}

func ApiDeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello Delete %s", name)
	}
}

func ApiPatchHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello Patch %s", name)
	}
}

func ApiHeadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello Head %s", name)
	}
}

func ApiOptionsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	}
}
