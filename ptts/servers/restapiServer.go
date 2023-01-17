package servers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	services "github.com/ptts-easy/ptts/services/RESTApi"
)

func RunApiServer(server_ip string, server_port string) {

	router := gin.Default()

	rest := router.Group("/rest")
	{
		rest.GET("/*name", services.ApiGetHandler())       //retrieve a record
		rest.POST("/*name", services.ApiPostHandler())     //create one
		rest.PUT("/*name", services.ApiPutHandler())       //update a record
		rest.DELETE("/*name", services.ApiDeleteHandler()) //delete one
		rest.PATCH("/*name", services.ApiPatchHandler())
		rest.HEAD("/*name", services.ApiHeadHandler())
		rest.OPTIONS("/*name", services.ApiOptionsHandler())
	}

	fmt.Println("RunApiServer::", server_port)

	router.Run(server_ip + ":" + server_port)
}
