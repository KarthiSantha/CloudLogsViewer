package handlers

import (
	"net/http"
	"strings"

	"github.com/KarthiSantha/auth/Service"
	"github.com/KarthiSantha/auth/model"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		log.Println("End of Logging Middleware")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("JWT Logging middle ware Starts ------>>>>>> SECURE ------->>> ")
		reqToken := r.Header.Get(model.JwtToken)
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		claims, err := Service.IsJwtTokenValid(reqToken)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r.Header.Set(model.UserIdentifier, claims.Email)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
