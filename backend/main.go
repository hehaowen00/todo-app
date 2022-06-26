package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"host.local/todo-app/backend/auth"
	"host.local/todo-app/backend/models"
	"host.local/todo-app/backend/routes"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := mysql.NewConfig()
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASSWORD")
	config.Net = "tcp"
	config.Addr = "db"
	config.DBName = os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	auth, err := auth.NewAuth()
	if err != nil {
		log.Fatal(err)
	}

	conn := models.NewConn(db)

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
