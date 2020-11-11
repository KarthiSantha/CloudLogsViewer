package handlers

import (
	"log"
	"net/http"

	//handlers

	"github.com/gorilla/mux"
)

func RegisterHandlers() {

	log.Print("Resitering all handlers to paths")
	r := mux.NewRouter() //.Schemes("http").Subrouter()
	//Define router Templates to be used
	post := r.Methods("POST").Subrouter()

	authRouter := post.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", Signup)
	authRouter.HandleFunc("/authenticate", Authenticate)

	http.ListenAndServe(":8080", r)

}
