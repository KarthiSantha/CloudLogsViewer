package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KarthiSantha/auth/Repository"
	"github.com/KarthiSantha/auth/model"
	log "github.com/sirupsen/logrus"
)

func UserProfile(rw http.ResponseWriter, req *http.Request) {

	email := req.Header.Get(model.UserIdentifier)

	userRepo := Repository.UserRepositoryMySQLImpl{}
	u, err := userRepo.GetByEmail(email)
	if err != nil {
		http.Error(rw, string(err.Error()), http.StatusInternalServerError)
		return

	}
	b, err := json.Marshal(u)
	rw.Write(b)

	log.Print("Hello Controller is called successfully by user ", email)
}
