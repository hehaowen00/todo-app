package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"host.local/todo-app/backend/auth"
	"host.local/todo-app/backend/models"
	"host.local/todo-app/backend/routes"
)

func main() {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal(err)
	}

	auth, err := auth.NewAuth()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := models.NewConn(db)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()

	routes.Init(&auth, conn)
	routes.RegisterRoutes(router)

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
