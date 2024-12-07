package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserType string

const (
	Admin UserType = "Admin"
	User  UserType = "User"
)

type registerInput struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Type     UserType `jons:"type"`
}

type loginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (app *application) register(c *gin.Context) {
	var input registerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": "true"})
}

func (app *application) login(c *gin.Context) {
	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": "true"})
}
