package model

import (
	"errors"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user User) IsValid() (bool, error) {

	return true, errors.New("can't work with 42")
}
