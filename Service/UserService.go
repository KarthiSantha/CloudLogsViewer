package Service

import (
	"errors"

	"github.com/KarthiSantha/auth/Repository"
	"github.com/KarthiSantha/auth/model"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	SignUp(user *model.User) (*model.User, error)
	Authenticate(login model.Login) []string
	IsDuplicate(user *model.User) (bool, error)
}

type UserServiceImpl struct {
}

func (userService UserServiceImpl) SignUp(user *model.User) (*model.User, error) {
	log.Print("User Trying to Signup is ", user)
	isValid, err := user.IsValid()
	if (!isValid) && (err != nil) {
		return nil, err
	}

	exists, err := userService.IsDuplicate(user)
	if err != nil {
		return nil, err
	}
	userRepo := Repository.UserRepositoryMySQLImpl{} //User Repo Creation
	if !exists {
		num, err := userRepo.Store(user)
		if err != nil {
			log.Print("Store User Repository method Failed ", err)
			return nil, err
		}
		if num == 1 {
			log.Print("User is On Boarded ", user)
		}
	} else {
		log.Print("Already Signed Up User Tried to SignUp ", user)
		return user, errors.New("You already have an Account " + user.Email)
	}

	return user, nil
}

func (userService UserServiceImpl) IsDuplicate(user *model.User) (bool, error) {
	log.Print("User Duplication Checking")
	userRepo := Repository.UserRepositoryMySQLImpl{} //User Repo Creation
	email := user.Email
	log.Print("Email is fetched" + email)
	u, err := userRepo.GetByEmail(email)
	log.Print("Fetched User by Email ", u)
	if err != nil {
		return false, err
	}
	if u.Email == "" {
		log.Print("User Is not Found")
		return false, nil
	}

	return true, nil

}

func (userService UserServiceImpl) SignIn(login *model.Login) (bool, error) {
	log.Print("User Login in Request ", login)
	email := login.Email

	userRepo := Repository.UserRepositoryMySQLImpl{} //User Repo Creation

	u, err := userRepo.GetByEmail(email)
	log.Print("Fetched User by Email ", u)
	if err != nil {
		return false, err
	}
	if u.Email == "" {
		log.Print("User not Found")
		return false, errors.New("Invalid User")
	}

	if u.Password != login.Password {
		log.Print("Passwords Dont Match")
		return false, errors.New("Invalid Password")
	}

	return true, nil
}
