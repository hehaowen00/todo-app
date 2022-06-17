package models

import (
	"context"
	"database/sql"
	"errors"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Migrate(c *sql.DB) error {
	_, err := c.Exec(`
  CREATE TABLE IF NOT EXISTS users (
    id integer PRIMARY KEY,
    username varchar(64) unique,
    password blob unique
  )`)

	return err
}

func (c *Conn) UserExists(user *User) (bool, error) {
	const CHECK_IF_USER_EXISTS = `
	SELECT username FROM users
	WHERE username = ?
	`

	stmt, err := c.db.Prepare(CHECK_IF_USER_EXISTS)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRow(user.Username)

	var username string

	err = row.Scan(&username)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, err
}

func (c *Conn) AddUser(user *User) error {
	ctx := context.Background()
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	const ADD_USER_QUERY = `
	INSERT INTO users (username, password)
	VALUES (?, ?)
	`

	stmt, err := tx.Prepare(ADD_USER_QUERY)
	if err != nil {
		return err
	}

	token, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = token

	res, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = id

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) VerifyUser(user *User) error {
	const GET_USER_QUERY = `
  SELECT id, password
  FROM users
  WHERE username = ?;
  `

	stmt, err := c.db.Prepare(GET_USER_QUERY)
	if err != nil {
		return err
	}

	var id int64
	var password string

	row := stmt.QueryRow(user.Username)

	err = row.Scan(&id, &password)
	if err != nil {
		return err
	}

	correct, err := VerifyPassword(user.Password, password)
	if err != nil {
		return err
	}

	if correct {
		user.Id = id
		user.Password = ""
		return nil
	}

	return errors.New("invalid credentials")
}
