package handlers

import (
	"database/sql"
	"log"
	"net/http"

	//handlers

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
	Port     string
}

func RegisterHandlers(Port string) {

	log.Print("Resitering all handlers to paths")
	r := mux.NewRouter() //.Schemes("http").Subrouter()
	//Define router Templates to be used
	r.Use(LoggingMiddleware)

	post := r.Methods("POST").Subrouter()

	authRouter := post.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", Signup)
	authRouter.HandleFunc("/authenticate", Authenticate)

	http.ListenAndServe(":"+Port, r)

}
