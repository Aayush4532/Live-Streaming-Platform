package controllers

import (
	"context"
	"net/http"
	"webinar/src/config"
	"webinar/src/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c *gin.Context) {
	type LoginInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	key := "host:" + input.Email

	storedHash, err := config.RDB.HGet(ctx, key, "password_hash").Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(storedHash),
		[]byte(input.Password),
	); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(input.Email, "HOST")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   7200,
		HttpOnly: true,
		Secure:   true,                 
		SameSite: http.SameSiteNoneMode, 
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}

func CreateHostController(c *gin.Context) {
	type CreateHostInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var input CreateHostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	key := "host:" + input.Email

	exists, err := config.RDB.Exists(ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
		return
	}
	if exists == 1 {
		c.JSON(http.StatusConflict, gin.H{"error": "host already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hash failed"})
		return
	}

	if err := config.RDB.HSet(ctx, key, map[string]interface{}{
		"password_hash": string(hash),
		"role":          "HOST",
		"active":        "true",
	}).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "host created successfully"})
}
