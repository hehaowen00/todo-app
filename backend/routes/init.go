package routes

import (
	"errors"

	"github.com/julienschmidt/httprouter"
	"host.local/todo-app/backend/auth"
	"host.local/todo-app/backend/models"
)

var (
	authInstance *auth.Auth
	dbConn       *models.Conn
)

func Init(auth *auth.Auth, conn *models.Conn) error {
	authInstance = auth
	if authInstance == nil {
		return errors.New("auth is nil")
	}

	dbConn = conn
	if dbConn == nil {
		return errors.New("database connection is nil")
	}

	return nil
}

func RegisterRoutes(router *httprouter.Router) {
	router.GET("/lists", requireAuth(GetLists))
	router.POST("/lists", requireAuth(AddList))
	router.PUT("/lists", requireAuth(UpdateList))
	router.DELETE("/lists", requireAuth(DeleteList))

	router.GET("/todos/:id", requireAuth(GetTodos))
	router.POST("/todos", requireAuth(AddTodo))
	router.PUT("/todos", requireAuth(UpdateTodo))
	router.DELETE("/todos", requireAuth(DeleteTodo))

	router.GET("/todo/:id", requireAuth(GetTodo))

	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/verify", requireAuth(VerifyToken))
}
