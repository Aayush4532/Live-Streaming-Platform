package routes

import (
	"webinar/src/controllers"
	"webinar/src/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes (r *gin.RouterGroup) {
	r.GET("/join-seminar", middleware.UserMiddleware(), controllers.StreamAccessController); // this route is for users to get access of the live stream,
	// this route will be called after each 3 minutes for refreshing the new sign token to make security standard.
}

func SetupAuthRoutes (r *gin.RouterGroup) {
	r.POST("/login", controllers.LoginController); // log in for only host
	r.POST("/register", middleware.HostMiddleware(), controllers.CreateHostController); // this route is used to register a new host for the application
}

func HostRoutes (r *gin.RouterGroup) {
	r.GET("/create", middleware.HostMiddleware(), controllers.CreateSeminarController); // host route for only for host to start the stream.
	r.GET("/check", middleware.HostMiddleware(), controllers.CheckHostController) // this is only for let frontend check if to give access of "/host" to user.
}


func MediaRoutes (r *gin.RouterGroup) {
	r.POST("/validate", controllers.ValidateSeminarController); // this route will only be called nginx to verify if the stream should be allowed to start or not.
}