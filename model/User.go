package model

import (
	"errors"
	"regexp"
)

var MIN_USERNAME_LENGTH int = 8
var MIN_PASSWORD_LENGTH int = 8
var MIN_NUM_SPECIAL_CHAR int = 1
var MIN_NUM_CAPS_CHAR int = 1
var MIN_NUM_NUMBERS int = 2
var IsSpaceAllowed bool = false
var AllowedSpecialChars []string = []string{"_", "@", "#", "&"}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user User) IsValid() (bool, error) {

	if len(user.Username) == 0 {
		return false, errors.New("Username cannot be Empty")
	}

	if len(user.Email) == 0 {
		return false, errors.New("Email cannot be Empty")
	}

	if len(user.Password) == 0 {
		return false, errors.New("Password cannot be Empty")
	}

	if len(user.Username) < MIN_USERNAME_LENGTH {
		return false, errors.New("UserName should have a minimum length of  " + string(MIN_USERNAME_LENGTH))
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	IsValidEmail := re.MatchString(user.Email)
	if !IsValidEmail {
		return false, errors.New("Inavlid Email ID " + user.Email)
	}

	if len(user.Password) < MIN_USERNAME_LENGTH {
		return false, errors.New("Password should have a minimum length of  " + string(MIN_USERNAME_LENGTH))
	}

	if len(user.Password) < MIN_USERNAME_LENGTH {
		return false, errors.New("Password should have a minimum length of  " + string(MIN_USERNAME_LENGTH))
	}

	return true, nil
}
