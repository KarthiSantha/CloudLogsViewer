package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/KarthiSantha/auth/Service"
	"github.com/KarthiSantha/auth/model"
)

func Signup(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var user model.User

	err := decoder.Decode(&user)
	if err != nil {
		b, _ := json.Marshal(err)
		http.Error(rw, string(b), http.StatusInternalServerError)
		return
	}

	userService := Service.UserServiceImpl{}

	u, err := userService.SignUp(&user)
	//log.Print("Outside if block User Service SignUp Error ", err)
	if err != nil {
		log.Print("User Service SignUp Error ", err)

		http.Error(rw, string(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Print("User sent is ", u)

	b, err := json.Marshal(user)

	rw.Write(b)

}

func Authenticate(rw http.ResponseWriter, req *http.Request) {
	log.Print("Authentication request has arrived ")

	decoder := json.NewDecoder(req.Body)
	var login model.Login

	err := decoder.Decode(&login)
	if err != nil {
		log.Print("Decoding of Login Failed ")
		b, _ := json.Marshal(err)
		http.Error(rw, string(b), http.StatusInternalServerError)
		return
	}

	userService := Service.UserServiceImpl{}

	IsAuthenticated, err := userService.SignIn(&login)
	if err != nil {
		log.Print("User Service SignIn Error ", err)

		rw.Write([]byte(err.Error()))
		return
	}

	if !IsAuthenticated {
		log.Print("User Authentication Failed")

		rw.Write([]byte("Login Failed"))
		return
	}
	log.Print("Authentication Success")
	token, r, err := Service.GetJwtToken(login.Email)

	if err != nil {
		log.Print("Login Success Token Generation Failed")
		rw.Write([]byte("Login Success Token Generation Failed"))
		return
	}
	refreshToken, err := Service.GetRefreshToken(login.Email, r)
	if err != nil {
		log.Print("Login Success Refresh Token Generation Failed")
		rw.Write([]byte("Login Success Refresh Token Generation Failed"))
		return
	}
	rw.Header().Set("Authorization", token)
	rw.Header().Set("RefreshToken", refreshToken)
	rw.Write([]byte("Login Success"))
}

func Token(rw http.ResponseWriter, req *http.Request) {
	log.Print("Token Refresh request has arrived ")

	reqToken := req.Header.Get(model.JwtToken)
	refreshToken := req.Header.Get(model.RefreshToken)
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	// Do stuff here

	isValid, err := Service.ValidateRefreshToken(reqToken, refreshToken)
	if !isValid || err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}
	claims, err := Service.IsJwtTokenValid(refreshToken)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}

	token, _, err := Service.GetJwtToken(claims.Email)
	rw.Header().Set(model.JwtToken, token)
	rw.Header().Set(model.RefreshToken, refreshToken)
	rw.Write([]byte("Token Refreshed"))

}
