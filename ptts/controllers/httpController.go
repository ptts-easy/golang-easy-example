package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "/http"
func HTTPHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "http.html", gin.H{
			"title":   "HTTP",
			"http":    "active",
			"content": "",
		})
	}
}

func HTTPInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		body, _ := io.ReadAll(c.Request.Body)

		get_param, _ := json.Marshal(c.Request.URL.Query())

		header, _ := json.Marshal(c.Request.Header)

		resp := map[string]string{"ip": c.Request.RemoteAddr,
			"method":     c.Request.Method,
			"url":        c.Request.URL.Path,
			"get_param":  string(get_param),
			"post_param": string(body),
			"header":     string(header),
			"body":       string(body)}

		c.JSON(http.StatusOK, gin.H{
			"ip":         resp["ip"],
			"method":     resp["method"],
			"url":        resp["url"],
			"get_param":  resp["get_param"],
			"post_param": resp["post_param"],
			"header":     resp["header"],
			"body":       resp["body"],
		})

	}
}
