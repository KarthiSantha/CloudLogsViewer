package main

import (
	"log"

	"auth/v1.0/handlers/registerHandlers"
)

func main() {
	registerHandlers.RegisterHandlers()

	log.Print("Main Method Started")

}
