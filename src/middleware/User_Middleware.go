package middleware

import "github.com/gin-gonic/gin"

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// function to get data from paf application for user verification.
		c.Next()
	}
}