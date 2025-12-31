package main

import (
	"log"
	"os"
	"time"

	"webinar/src/config"
	"webinar/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using platform env vars")
	}

	log.Println("BOOTING APPLICATION")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://live-streaming-platform-sigma.vercel.app",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		AllowCredentials: true, 
		MaxAge:           12 * time.Hour,
	}))

	userGroup := r.Group("/api/user")
	authGroup := r.Group("/api/auth")
	hostGroup := r.Group("/api/host")
	mediaServerGroup := r.Group("/api/media")

	routes.SetupAuthRoutes(authGroup)
	routes.UserRoutes(userGroup)
	routes.HostRoutes(hostGroup)
	routes.MediaRoutes(mediaServerGroup)

	if err := config.InitRedis(); err != nil {
		log.Println("Redis not connected, continuing without Redis:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on port", port)
	r.Run(":" + port)
}
