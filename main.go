package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/router"
)

func init() {
	//initializer.ConnectEnv()
}

func main() {
	app := gin.Default()

	router.ApiRouter(app)

	app.Run("localhost:3000")
}
