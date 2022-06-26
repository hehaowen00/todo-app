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
	router.GET("/api/lists", requireAuth(getLists))
	router.GET("/api/lists/:id", requireAuth(getList))
	router.POST("/api/lists", requireAuth(addList))
	router.PUT("/api/lists", requireAuth(updateList))
	router.DELETE("/api/lists", requireAuth(deleteList))

	router.GET("/api/todos/:id", requireAuth(getTodos))
	router.POST("/api/todos", requireAuth(addTodo))
	router.PUT("/api/todos", requireAuth(updateTodo))
	router.DELETE("/api/todos", requireAuth(deleteTodo))

	router.GET("/api/todo/:id", requireAuth(getTodo))

	router.POST("/api/register", register)
	router.POST("/api/login", login)
	router.POST("/api/verify", requireAuth(verifyToken))

	router.GET("/api/user/data", requireAuth(exportUserData))
	router.PUT("/api/user/profile", requireAuth(updateUsername))
	router.PUT("/api/user/password", requireAuth(updatePassword))
	router.DELETE("/api/user", requireAuth(deleteUser))
}
