package main

import (
	"metaverse/http/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) register(c *gin.Context) {
	var input handlers.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": "true"})
}

func (app *application) login(c *gin.Context) {
	var input handlers.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": "true"})
}
