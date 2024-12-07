package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "available", "version": version, "port": app.config.port})
}
