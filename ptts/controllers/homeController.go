package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// "/" => "hello"
func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Redirect(http.StatusMovedPermanently, "/hello")
	}
}

// "/template"
func TemplateGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		now := time.Now() // current local time
		timestamp := now.Unix()

		type MyObj struct {
			User string
			Msg  string
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":         "Template",
			"template":      "active",
			"content":       "Hello World ...",
			"object":        MyObj{User: "user001", Msg: "msg001"},
			"timestamp":     timestamp,
			"timestamp_odd": timestamp%2 == 1,
			"users":         []string{"user1", "user2", "user3"},
		})
	}
}

// "/route"
func RouteDefaultGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "route.html", gin.H{
			"title":   "Route",
			"route":   "active",
			"name":    "",
			"action":  "",
			"act1":    "active",
			"message": "This is default route",
			"content": "",
		})
	}
}

func RouteGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.Param("name")
		action := c.Param("action")

		if len(action) > 0 && action[0] == '/' {
			action = action[1:]
		}

		message := name + " is " + action
		data := gin.H{
			"title":   "Route",
			"route":   "active",
			"name":    name,
			"action":  action,
			action:    "active",
			"message": message,
			"content": "",
		}

		c.HTML(http.StatusOK, "route.html", data)
	}
}
