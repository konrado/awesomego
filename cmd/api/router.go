package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) router() *gin.Engine {
	r := gin.Default()

	r.GET("/healthcheck", app.healthcheckHandler)
	r.POST("/authors", app.createAuthor)
	r.GET("/authors/:id", app.viewAuthor)

	return r
}
