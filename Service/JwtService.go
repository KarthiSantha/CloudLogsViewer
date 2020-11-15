package Service

import (
	"errors"
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

//github.com/dgrijalva/jwt-go v3.0.0+incompatible

const jwt_expiry = 10
const rt_expiry = 10 * 24 * 60

// Create the JWT key used to create the signature
var JwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
	RefreshNumber *rand.Rand `json:"refresh_random_number"`
}

func GetJwtToken(email string) (string, *rand.Rand, error) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	log.Print(r)

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(jwt_expiry * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
		RefreshNumber: r,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", nil, err
	}
	return tokenString, r, nil
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

func GetRefreshToken(email string, r *rand.Rand) (string, error) {

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(rt_expiry * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
		RefreshNumber: r,
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

func ValidateRefreshToken(jwtToken string, refreshToken string) (bool, error) {

	// Initialize a new instance of `Claims`
	claimsRefreshToken := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	refreshtkn, err := jwt.ParseWithClaims(refreshToken, claimsRefreshToken, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return false, err
	}

	if !refreshtkn.Valid {
		return false, errors.New("Refresh Token is not valid")
	}

	return true, nil
}
