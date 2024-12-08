package main

import (
	"metaverse/http/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) register(c *gin.Context) {
	var input data.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userModel := data.NewUserModel(app.db)

	err := userModel.Create(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": true})
}

func (app *application) login(c *gin.Context) {
	var input data.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userModel := data.NewUserModel(app.db)

	user, err := userModel.CheckUser(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": user.ID})
}
