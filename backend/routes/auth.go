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
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Could not decode JSON")
		return
	}

	if user.Username == "" || user.Password == "" {
		log.Println("invalid request body")
		jsonMessage(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	user.Trim()

	exists, err := dbConn.UsernameExists(&user)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Server Error")
		return
	}

	if exists {
		jsonMessage(w, http.StatusBadRequest, "User already exists")
		return
	}

	err = dbConn.AddUser(&user)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to add user")
		return
	}

	jsonMessage(w, http.StatusOK, "User added")
}

func Login(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	user.Trim()

	err = dbConn.VerifyUser(&user)

	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusUnauthorized, "Unauthorized User")
		return
	}

	token, err := authInstance.CreateJWT(user.Id, user.Username)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to create token")
		return
	}

	jsonMessage(w, http.StatusOK, token)
}

func VerifyToken(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	username := req.Context().Value(usernameKey).(string)

	token, err := authInstance.CreateJWT(userID, username)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to create token")
		return
	}

	jsonMessage(w, http.StatusOK, token)
}

func requireAuth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, hs httprouter.Params) {
		header := req.Header.Get("Authorization")
		if header == "" {
			log.Println("Missing Header")
			w.Header().Set("WWW-Authenticate", "Bearer realm=restricted")
			jsonMessage(w, http.StatusUnauthorized, "Invalid auth token")
			return
		}

		parts := strings.Split(header, "Bearer ")
		if len(parts) != 2 {
			log.Println("Failed to parse header")
			w.Header().Set("WWW-Authenticate", "Bearer realm=restricted")
			jsonMessage(w, http.StatusUnauthorized, "Invalid auth token")
			return
		}

		// validate token
		claims, err := authInstance.ValidateJWT(parts[1])
		if err != nil {
			log.Println(err)
			jsonMessage(w, http.StatusUnauthorized, "Invalid auth token")
			return
		}

		var user models.User
		user.Id = claims.Id
		user.Username = claims.Username

		user.Trim()

		exists, err := dbConn.UserExists(&user)
		if err != nil {
			log.Println(err)
			jsonMessage(w, http.StatusInternalServerError, "Failed to validate token")
			return
		}

		if !exists {
			log.Println("User does not exist:", user.Username)
			jsonMessage(w, http.StatusUnauthorized, "Invalid auth token")
			return
		}

		ctx := context.WithValue(req.Context(), idKey, user.Id)
		ctx = context.WithValue(ctx, usernameKey, user.Username)

		handle(w, req.WithContext(ctx), hs)
	}
}
