package routes

import (
	"webinar/src/controllers"
	"webinar/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupJoinRoutes(r *gin.RouterGroup) {
	r.POST("/:id", );
}

func SetupAuthRoutes (r *gin.RouterGroup) {
	r.POST("/login", controllers.LoginController); // log in the only for host
	r.POST("/register", middleware.HostMiddleware(), controllers.CreateHostController); // this route is used to register a new host for the application
}