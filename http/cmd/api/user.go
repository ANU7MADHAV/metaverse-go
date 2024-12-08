package main

import (
	"metaverse/http/internal/data"
	"metaverse/http/utils"
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

	token, err := utils.CreateToken(input.Username, string(input.Type))

	if err != nil {
		app.logger.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
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

	token, err := utils.CreateToken(user.Username, string(user.Type))

	if err != nil {
		app.logger.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
