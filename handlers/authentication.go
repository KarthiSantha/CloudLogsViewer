package handlers

import (
	"log"
	"net/http"
)

func Signup(rw http.ResponseWriter, req *http.Request) {
	log.Print("Signup API Initiated")

}

func Authenticate(rw http.ResponseWriter, req *http.Request) {
	log.Print("Authentication request has arrived ")

}
