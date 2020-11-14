package Repository

import (
	"database/sql"
	"fmt"

	"github.com/KarthiSantha/auth/model"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var DatabaseConnection *sql.DB

func CreateDatabase(configuration model.Config) error {
	log.Print("Database COnfig is ", configuration)
	serverName := configuration.DBHostname + ":" + configuration.DBPort
	user := configuration.DBUsername
	password := configuration.DBPassword
	dbName := configuration.DBDatabase
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	DatabaseConnection, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Database Connection Failed")
		return err
	} else {
		log.Print("Database Connection Success")
	}

	return nil
}
