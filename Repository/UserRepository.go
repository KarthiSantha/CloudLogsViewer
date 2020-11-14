package Repository

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarthiSantha/auth/model"
)

type UserRepository interface {
	GetByEmail(email string) (*model.User, error)
	Store(user *model.User) (int64, error)
}

type RepositoryMySQL struct {
}

type UserRepositoryMySQLImpl struct {
}

func (userRepo UserRepositoryMySQLImpl) Store(user *model.User) (int64, error) {
	log.Print("User STore Precodure Called for user", user)

	databaseConnection := DatabaseConnection

	insForm, err := databaseConnection.Prepare("INSERT INTO user(username, email, password) VALUES(?,?,?)")
	if err != nil {
		log.Print("Database User Insert Query failed ", err)
		return 0, err
	}
	result, err := insForm.Exec(user.Username, user.Email, user.Password)
	num, err := result.RowsAffected()
	log.Print("Number of Rows Affected ", num)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (userRepo UserRepositoryMySQLImpl) GetByEmail(email string) (*model.User, error) {
	log.Print("Get User BY Email ----------------- ")
	databaseConnection := DatabaseConnection

	selDB, err := databaseConnection.Query("SELECT username,email,password FROM user WHERE email=?", email)
	if err != nil {
		log.Print("Database User GetByEmail Query failed ", err)
		return nil, err
	}
	var user model.User
	for selDB.Next() {

		var username, email, password string
		err = selDB.Scan(&username, &email, &password)
		if err != nil {
			panic(err.Error())
		}
		user.Username = username
		user.Email = email
		user.Password = password
	}
	return &user, nil

}
