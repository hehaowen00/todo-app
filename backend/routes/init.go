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
	router.GET("/api/lists", requireAuth(GetLists))
	router.POST("/api/lists", requireAuth(AddList))
	router.PUT("/api/lists", requireAuth(UpdateList))
	router.DELETE("/api/lists", requireAuth(DeleteList))

	router.GET("/api/todos/:id", requireAuth(GetTodos))
	router.POST("/api/todos", requireAuth(AddTodo))
	router.PUT("/api/todos", requireAuth(UpdateTodo))
	router.DELETE("/api/todos", requireAuth(DeleteTodo))

	router.GET("/api/todo/:id", requireAuth(GetTodo))

	router.POST("/api/register", Register)
	router.POST("/api/login", Login)
	router.POST("/api/verify", requireAuth(VerifyToken))
}
