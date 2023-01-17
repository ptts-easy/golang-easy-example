package controllers

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "github.com/ptts-easy/ptts/globals"
)

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		useremail := session.Get(globals.Userkey)

		if useremail != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"title":    "Login",
					"login":    "active",
					"username": globals.Username[fmt.Sprintf("%v", useremail)],
					"content":  "Please logout first",
				})
		} else {
			c.HTML(http.StatusOK, "login.html",
				gin.H{
					"title": "Login",
					"login": "active",
				})
		}
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		useremail := session.Get(globals.Userkey)

		if useremail != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"title":    "Login",
					"login":    "active",
					"content":  "Please logout first",
					"username": globals.Username[fmt.Sprintf("%v", useremail)],
				})
		} else {
			useremail := c.PostForm("email")
			password := c.PostForm("password")

			if EmptyUserPass(useremail, password) {
				c.HTML(http.StatusBadRequest, "login.html",
					gin.H{
						"title":   "Login",
						"login":   "active",
						"content": "Parameters can't be empty",
					})
			} else if !CheckUserPass(useremail, password) {
				c.HTML(http.StatusBadRequest, "login.html",
					gin.H{
						"title":   "Login",
						"login":   "active",
						"content": "Incorrect email or password",
					})
			} else {
				session.Set(globals.Userkey, useremail)
				if err := session.Save(); err != nil {
					c.HTML(http.StatusInternalServerError, "login.html",
						gin.H{
							"title":   "Login",
							"login":   "active",
							"content": "Failed to save session",
						})
				} else {
					c.Redirect(http.StatusFound, "/dashboard")
				}
			}
		}
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		useremail := session.Get(globals.Userkey)

		fmt.Println("logging out user:", useremail)

		if useremail == nil {
			fmt.Println("Invalid session token")
		} else {
			session.Delete(globals.Userkey)

			if err := session.Save(); err != nil {
				fmt.Println("Failed to save session:", err)
			}
		}

		c.Redirect(http.StatusFound, "/login")
	}
}

func CheckUserPass(useremail, password string) bool {

	log.Println("checkUserPass", useremail, password, globals.Userpass[useremail])

	if val, ok := globals.Userpass[useremail]; ok {
		log.Println(val, ok)
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func EmptyUserPass(useremail, password string) bool {
	return strings.Trim(useremail, " ") == "" || strings.Trim(password, " ") == ""
}
