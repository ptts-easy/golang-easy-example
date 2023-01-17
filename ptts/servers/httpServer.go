package servers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	globals "github.com/ptts-easy/ptts/globals"
	middleware "github.com/ptts-easy/ptts/middleware"
	routes "github.com/ptts-easy/ptts/routes"
)

func RunHttpServer(server_ip string, server_port string) {

	//	router := gin.Default()
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	router.Static("/assets", "./public/assets")
	router.Static("/libs", "./public/libs")
	router.StaticFile("/hello", "./public/hello.html")

	router.LoadHTMLGlob("./ptts/views/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired())
	routes.PrivateRoutes(private)

	fmt.Println("RunHttpServer::", server_port)

	router.Run(server_ip + ":" + server_port)
}
