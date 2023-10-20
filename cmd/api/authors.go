package main

import (
	"awesome.itkon.pl/internal/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"time"
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
	authorId, _ := uuid.Parse(c.Params.ByName("id"))

	author := data.Author{
		Id:        authorId,
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
	}

	extra := map[string]string{"version": version}

	c.JSON(http.StatusOK, data.View{Data: author, Extra: extra})
}
