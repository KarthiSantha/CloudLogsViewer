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
	r := mux.NewRouter()
	r.Use(LoggingMiddleware)

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", Signup).Methods("POST")
	authRouter.HandleFunc("/authenticate", Authenticate).Methods("POST")

	securedRouter := r.PathPrefix("/api").Subrouter()
	securedRouter.Use(JwtAuthMiddleware)

	securedRouter.HandleFunc("/token", Token).Methods("POST")
	securedRouter.HandleFunc("/getUserProfile", UserProfile).Methods("GET")

	http.ListenAndServe(":"+Port, r)

}
