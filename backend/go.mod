module host.local/todo-app/backend

go 1.18

replace host.local/todo-app/backend/auth => ./auth

replace host.local/todo-app/backend/models => ./models

replace host.local/todo-app/backend/routes => ./routes

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/mattn/go-sqlite3 v1.14.13
	github.com/rs/cors v1.8.2
	host.local/todo-app/backend/auth v0.0.0-00010101000000-000000000000
	host.local/todo-app/backend/models v1.0.0
	host.local/todo-app/backend/routes v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)
