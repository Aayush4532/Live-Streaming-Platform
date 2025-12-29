package main

import (
	"os"
	"webinar/src/config"
	"webinar/src/controllers"
	"webinar/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	joinGroup := r.Group("/api/join")
	authGroup := r.Group("/api/auth")
	routes.SetupAuthRoutes(authGroup)
	routes.SetupJoinRoutes(joinGroup)
	err := config.InitRedis()
	if err != nil {
		panic(err)
	}

	r.GET("/api/ws", controllers.WsHandler)

	r.Run(":" + os.Getenv("PORT"))
}
