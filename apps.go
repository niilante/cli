package arukas

import (
	"errors"
	"time"
)

type App struct {
	ID          string     `json:"-"`
	Name        string     `json:"name"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"-"`
	ContainerId string     `json:"-"`
	Container   *Container `json:"-"`
	User        *User      `json:"-"`
}

func (a App) GetID() string {
	return string(a.ID)
}

func (a *App) SetID(ID string) error {
	a.ID = ID
	return nil
}

func (a *App) SetToOneReferenceID(name, ID string) error {
	if name == "container" {
		if ID == "" {
			a.Container = nil
		} else {
			a.Container = &Container{ID: ID}
		}

		return nil
	} else if name == "user" {
		if ID == "" {
			a.User = nil
		} else {
			a.User = &User{ID: ID}
		}

		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}
