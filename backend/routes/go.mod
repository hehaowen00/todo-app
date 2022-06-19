module host.local/todo-app/backend/routes

go 1.18

replace host.local/todo-app/backend/models => ../models

replace host.local/todo-app/backend/auth => ../auth

require (
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	host.local/todo-app/backend/auth v0.0.0-00010101000000-000000000000 // indirect
	host.local/todo-app/backend/models v0.0.0-00010101000000-000000000000 // indirect
)
