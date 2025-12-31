package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func StreamAccessController(c *gin.Context) {
	baseURL := os.Getenv("STREAM_URI")
	secret := os.Getenv("STREAM_TOKEN_SIGN_URI")

	if baseURL == "" || secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server misconfigured"})
		return
	}

	path := c.DefaultQuery("path", "/hls/stream.m3u8")
	if !strings.HasPrefix(path, "/hls/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid path"})
		return
	}

	expiry := time.Now().Add(3 * time.Minute).Unix()
	expStr := strconv.FormatInt(expiry, 10)

	stringToSign := path + expStr

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	signedURL := baseURL + path + "?exp=" + expStr + "&sig=" + signature

	c.JSON(http.StatusOK, gin.H{
		"playback": gin.H{
			"url":        signedURL,
			"expires_at": expiry,
		},
		"refresh_in": 120,
	})
}
