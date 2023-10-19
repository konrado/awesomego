package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

var authors = map[int]string{
	1: "John Doe",
	2: "Johny Bravo",
	3: "Hector Alfa",
}

func (app *application) createAuthor(c *gin.Context) {
	r1 := rand.Intn(20)
	authors[r1] = "Added"
	c.Header("Location", fmt.Sprintf("/authors/%d", r1))
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
