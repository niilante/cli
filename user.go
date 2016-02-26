package arukas

import (
	// "errors"
	// "fmt"
	// "github.com/codegangsta/cli"
	// "os"
	"time"
)

type User struct {
	ID          string    `json:"-"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Provider    string    `json:"provider"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ConfirmedAt time.Time `json:"-"`
}

func (u User) GetID() string {
	return string(u.ID)
}

func (u *User) SetID(ID string) error {
	u.ID = ID
	return nil
}
