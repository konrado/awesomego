package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) healthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "available",
		"version":     fmt.Sprintf("%s", version),
		"environment": fmt.Sprintf("%s", app.config.env),
	})
}
