package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var authors = map[int]string{
	1: "John Doe",
	2: "Johny Bravo",
	3: "Hector Alfa",
}

func (app *application) createAuthor(c *gin.Context) {
	c.Header("Location", fmt.Sprintf("/authors/%d", 100))
	c.JSON(http.StatusCreated, gin.H{})
}

func (app *application) viewAuthor(c *gin.Context) {
	authorId, _ := strconv.Atoi(c.Params.ByName("id"))
	author, ok := authors[authorId]

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   authorId,
		"name": author,
	})
}
