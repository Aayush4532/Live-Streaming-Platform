package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"webinar/src/config"
	"webinar/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateSeminarController(c *gin.Context) { // for host to create the live streaming.
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	uniqueHex, err := utils.GenerateStreamKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate stream key",
		})
		return
	}

	streamKey := "paf_live_" + uniqueHex
	redisKey := "streamkey:" + streamKey

	err = config.RDB.Set(ctx, redisKey, "1", 5*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "redis error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"rtmp_uri":   os.Getenv("RTMP_URI"),
		"stream_key": streamKey,
	})
}

func ValidateSeminarController(c *gin.Context) { // this route is only for nginx
	streamKey := c.PostForm("name");   // RTMP stream key comes as form field "name"

	if streamKey == "" {
		c.Status(http.StatusForbidden)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	redisKey := "streamkey:" + streamKey

	exists, err := config.RDB.Exists(ctx, redisKey).Result()
	if err != nil || exists == 0 {
		c.Status(http.StatusForbidden)
		return
	}

	c.Status(http.StatusOK)
}

func CheckHostController (c *gin.Context) {
	c.JSON(200, gin.H {
		"message" : "authorised",
	})
}