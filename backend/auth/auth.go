package auth

import (
	"errors"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Id       int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Auth struct {
	secret []byte
}

func NewAuth() (Auth, error) {
	var auth Auth
	secret := make([]byte, 1024)

	_, err := rand.Read(secret)
	if err != nil {
		return auth, err
	}

	auth.secret = secret
	return auth, nil
}

func (a *Auth) CreateJWT(id int64, username string) (string, error) {
	expiration := time.Now().UTC().Add(24 * time.Hour)

	claims := Token{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(a.secret)
}

func (a *Auth) ValidateJWT(signedToken string) (*Token, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Token{},
		func(token *jwt.Token) (interface{}, error) {
			return a.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Token)
	if !ok {
		err = errors.New("invalid token")
		return nil, err
	}

	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		err = errors.New("token expired")
		return nil, err
	}

	return claims, nil
}
