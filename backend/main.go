package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"host.local/todo-app/backend/models"
)

type UserIDKey struct{}
type UsernameKey struct{}

var idKey = UserIDKey{}
var usernameKey = UsernameKey{}

var auth *Auth
var conn *models.Conn

func GetTodos(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	listID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	items, err := conn.GetTodoItems(userID, listID)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(items)
}

func GetTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	itemID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	var item models.TodoItem
	item.Id = itemID
	item.UserId = userID

	err = conn.GetTodoItem(&item)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(item)
}

func AddTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	item.UserId = userID
	item.Desc = strings.TrimSpace(item.Desc)
	if item.Desc == "" {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	err = conn.AddTodoItem(&item)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func UpdateTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	item.UserId = userID
	err = conn.UpdateTodoItem(&item)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func DeleteTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	item.UserId = userID
	err = conn.DeleteTodoItem(&item)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetLists(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	lists, err := conn.GetTodoLists(userID)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}

func AddList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	list.UserId = userID

	exists, err := conn.ListExists(&list)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("list with that name already exists")
		return
	}

	list.UserId = userID
	err = conn.AddTodoList(&list)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func UpdateList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	list.UserId = userID
	err = conn.UpdateTodoList(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func DeleteList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	list.UserId = userID
	err = conn.DeleteTodoList(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	err = conn.DeleteTodoItemsFromList(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	httpMessage(w, http.StatusOK)
}

func Register(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	exists, err := conn.UserExists(&user)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("user already exists")
		return
	}

	err = conn.AddUser(&user)
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

	err = conn.VerifyUser(&user)

	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusUnauthorized)
		return
	}

	token, err := auth.createJWT(user.Id, user.Username)
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

	token, err := auth.createJWT(userID, username)
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
		claims, err := auth.validateJWT(parts[1])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("failed to validate token")
			return
		}

		var id int64 = claims.Id
		var username string = claims.Username

		ctx := context.WithValue(req.Context(), idKey, id)
		ctx = context.WithValue(ctx, usernameKey, username)

		handle(w, req.WithContext(ctx), hs)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal(err)
	}

	auth_, err := NewAuth()
	if err != nil {
		log.Fatal(err)
	}

	auth = &auth_
	conn = models.NewConn(db)

	err = conn.Migrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Migrate(&models.TodoList{})
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Migrate(&models.TodoItem{})
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()

	router.GET("/todos/:id", requireAuth(GetTodos))
	router.GET("/todo/:id", requireAuth(GetTodo))
	router.POST("/todos", requireAuth(AddTodo))
	router.PUT("/todos", requireAuth(UpdateTodo))
	router.DELETE("/todos", requireAuth(DeleteTodo))

	router.GET("/lists", requireAuth(GetLists))
	router.POST("/lists", requireAuth(AddList))
	router.PUT("/lists", requireAuth(UpdateList))
	router.DELETE("/lists", requireAuth(DeleteList))

	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/verify", requireAuth(VerifyToken))

	options := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		Debug:            false,
	}

	handler := cors.New(options).Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
