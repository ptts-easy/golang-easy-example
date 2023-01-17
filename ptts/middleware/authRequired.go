package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ptts-easy/ptts/globals"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)

		useremail := session.Get(globals.Userkey)

		if useremail == nil {
			fmt.Println("User not logged in")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
