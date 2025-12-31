package main

import (
	"os"
	"webinar/src/config"
	"webinar/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	userGroup := r.Group("/api/user")
	authGroup := r.Group("/api/auth")
	hostGroup := r.Group("/api/host");
	mediaServerGroup := r.Group("/api/media");
	routes.SetupAuthRoutes(authGroup)
	routes.UserRoutes(userGroup)
	routes.HostRoutes(hostGroup);
	routes.MediaRoutes(mediaServerGroup);
	
	err := config.InitRedis()
	if err != nil {
		panic(err)
	}

	r.Run(":" + os.Getenv("PORT"))
}
