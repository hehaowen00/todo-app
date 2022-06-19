package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"host.local/todo-app/backend/models"
)

type UserIDKey struct{}
type UsernameKey struct{}

var (
	idKey       = UserIDKey{}
	usernameKey = UsernameKey{}
)

func Register(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	user.Trim()

	exists, err := dbConn.UserExists(&user)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("user already exists")
		return
	}

	err = dbConn.AddUser(&user)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	httpMessage(w, http.StatusOK)
}

func Login(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	user.Trim()

	err = dbConn.VerifyUser(&user)

	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusUnauthorized)
		return
	}

	token, err := authInstance.CreateJWT(user.Id, user.Username)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func VerifyToken(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	username := req.Context().Value(usernameKey).(string)

	token, err := authInstance.CreateJWT(userID, username)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func requireAuth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, hs httprouter.Params) {
		header := req.Header.Get("Authorization")
		if header == "" {
			log.Println("Missing Header")
			w.Header().Set("WWW-Authenticate", "Bearer realm=restricted")
			httpMessage(w, http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, "Bearer ")
		if len(parts) != 2 {
			log.Println("Failed to parse header")
			w.Header().Set("WWW-Authenticate", "Bearer realm=restricted")
			httpMessage(w, http.StatusUnauthorized)
			return
		}

		// validate token
		claims, err := authInstance.ValidateJWT(parts[1])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("failed to validate token")
			return
		}

		var id int64 = claims.Id
		var username string = claims.Username

		var user models.User
		user.Username = username

		user.Trim()

		exists, err := dbConn.UserExists(&user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("failed to validate token")
			return
		}

		if !exists {
			log.Println("user does not exist:", username)
			httpMessage(w, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), idKey, id)
		ctx = context.WithValue(ctx, usernameKey, username)

		handle(w, req.WithContext(ctx), hs)
	}
}
