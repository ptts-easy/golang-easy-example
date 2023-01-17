package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "github.com/ptts-easy/ptts/globals"
)

// "/dashboard"
func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		useremail := session.Get(globals.Userkey)

		if useremail == nil {
			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"title":     "Dashboard",
				"dashboard": "active",
				"noneuser":  true,
				"content":   "This is a dashboard",
			})
		} else {
			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"title":     "Dashboard",
				"dashboard": "active",
				"username":  globals.Username[fmt.Sprintf("%v", useremail)],
				"content":   "This is a dashboard",
			})

		}
	}
}
