package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	repo "github.com/KarthiSantha/auth/Repository"
	"github.com/KarthiSantha/auth/Service"
	"github.com/KarthiSantha/auth/handlers"
)

func init() {
	env := flag.String("env", "dev", "Environment dev/QA/property to choose config file appropriately")
	flag.Parse()
	config := Service.InitializeConfig(*env)
	Service.JwtKey = []byte(config.JWTSecretKey)

	err := repo.CreateDatabase(config)
	if err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	}

	log.Print("Database COnnection is ", repo.DatabaseConnection.Stats())

	err = repo.DatabaseConnection.Ping()
	if err != nil {
		panic(err)
	}

	Port := config.Port
	handlers.RegisterHandlers(Port)
}

func main() {

	log.Print("Main Method Started")

}
