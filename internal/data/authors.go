package data

import (
	"github.com/google/uuid"
	"time"
)

type Author struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"-"`
}
