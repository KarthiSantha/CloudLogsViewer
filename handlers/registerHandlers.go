package handlers

import (
	"database/sql"
	"net/http"

	log "github.com/sirupsen/logrus"

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

	//post := r.Methods("POST").Subrouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", Signup).Methods("POST")
	authRouter.HandleFunc("/authenticate", Authenticate).Methods("POST")

	securedRouter := r.PathPrefix("/Secure").Subrouter()
	securedRouter.Use(JwtAuthMiddleware)

	securedRouter.HandleFunc("/token", Token).Methods("POST")
	securedRouter.HandleFunc("/getUserProfile", UserProfile).Methods("GET")

	http.ListenAndServe(":"+Port, r)

}
