package ptts

import (
	"fmt"
	"io"
	"log"
	"os"

	gin "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	servers "github.com/ptts-easy/ptts/servers"
)

func Launcher() {
	f, _ := os.Create(os.Getenv("SERVER_LOG"))
	gin.DefaultWriter = io.MultiWriter(f)

	fmt.Println("launcher")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SERVER_IP := os.Getenv("SERVER_IP")
	HTTP_SERVER_PORT := os.Getenv("HTTP_SERVER_PORT")
	API_SERVER_PORT := os.Getenv("API_SERVER_PORT")
	FILE_SERVER_PORT := os.Getenv("FILE_SERVER_PORT")

	fmt.Println("HTTP_SERVER_PORT = ", HTTP_SERVER_PORT)
	fmt.Println("API_SERVER_PORT = ", API_SERVER_PORT)
	fmt.Println("FILE_SERVER_PORT = ", FILE_SERVER_PORT)

	go servers.RunApiServer(SERVER_IP, API_SERVER_PORT)

	go servers.RunUploadServer(SERVER_IP, FILE_SERVER_PORT)

	servers.RunHttpServer(SERVER_IP, HTTP_SERVER_PORT)
}
