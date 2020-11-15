package Service

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//github.com/dgrijalva/jwt-go v3.0.0+incompatible

const expiry = 10

// Create the JWT key used to create the signature
var JwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GetJwtToken(email string) (string, error) {

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(expiry * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}
	return tokenString, nil
}

func IsJwtTokenValid(tknStr string) (*Claims, error) {

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("Token is not valid")
	}
	return claims, nil
}
