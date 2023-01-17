package servers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	services "github.com/ptts-easy/ptts/services/filestorage"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("CORSMiddleware --- OPTIONS")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}

func RunUploadServer(server_ip string, server_port string) {

	UPLOAD_PATH := os.Getenv("UPLOAD_PATH")

	router := gin.Default()

	router.Use(CORSMiddleware())

	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:38080"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	*/

	// Set a lower memory limit for multipart forms (default is 32 MiB)

	router.MaxMultipartMemory = 1024 << 20 // 1024 MiB

	//	router.StaticFS("/storage", http.Dir(UPLOAD_PATH))

	storage_path, _ := filepath.Abs(UPLOAD_PATH)

	router.POST("/upload_single", services.ServiceUploadSingleFile(storage_path))
	router.POST("/upload_multi", services.ServiceUploadMultiFile(storage_path))
	router.POST("/show_storage", services.ServiceShowStorage(storage_path))
	router.POST("/clear_storage", services.ServiceClearStorage(storage_path))

	fmt.Println("RunUploadServer::", server_port)

	router.Run(server_ip + ":" + server_port)
}
