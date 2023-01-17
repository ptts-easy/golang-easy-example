package routes

import (
	"github.com/gin-gonic/gin"

	controllers "github.com/ptts-easy/ptts/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/", controllers.IndexGetHandler())
	g.GET("/template", controllers.TemplateGetHandler())
	g.GET("/route", controllers.RouteDefaultGetHandler())
	g.GET("/route/:name/*action", controllers.RouteGetHandler())
	g.GET("/fileupload", controllers.FileuploadGetHandler())

	g_http := g.Group("/http")
	{
		g_http.GET("/", controllers.HTTPHandler())
		g_http.GET("/info", controllers.HTTPInfoHandler())
		g_http.POST("/info", controllers.HTTPInfoHandler())
		g_http.PUT("/info", controllers.HTTPInfoHandler())
		g_http.DELETE("/info", controllers.HTTPInfoHandler())
		g_http.PATCH("/info", controllers.HTTPInfoHandler())
		g_http.HEAD("/info", controllers.HTTPInfoHandler())
		g_http.OPTIONS("/info", controllers.HTTPInfoHandler())
	}

	g.GET("/restapi", controllers.RestApiGetHandler())

	g_api := g.Group("/rest")
	{

		g_api.Any("/*name", controllers.RestApiMsgHandler())
	}

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.Any("/logout", controllers.LogoutGetHandler())
	g.GET("/dashboard", controllers.DashboardGetHandler())
}
