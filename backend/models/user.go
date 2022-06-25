package models

import (
	"context"
	"database/sql"
	"errors"
	"strings"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Trim() {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
}

func (c *Conn) UsernameExists(username string) (bool, error) {
	const CHECK_QUERY = `
	SELECT username FROM users
	WHERE username = ?
	`

	stmt, err := c.db.Prepare(CHECK_QUERY)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRow(username)

	var result string

	err = row.Scan(&result)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, err
}

func (c *Conn) UserExists(user *User) (bool, error) {
	const CHECK_QUERY = `
	SELECT username FROM users
	WHERE id = ? AND username = ?
	`

	stmt, err := c.db.Prepare(CHECK_QUERY)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRow(user.Id, user.Username)

	var username string

	err = row.Scan(&username)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, err
}

func (c *Conn) AddUser(user *User) error {
	const INSERT_QUERY = `
	INSERT INTO users (username, password)
	VALUES (?, ?)
	`

	ctx := context.Background()

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(INSERT_QUERY)
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

	return err
}

func (c *Conn) VerifyUser(user *User) error {
	const GET_QUERY = `
	SELECT id, password
	FROM users
	WHERE username = ?;
	`

	stmt, err := c.db.Prepare(GET_QUERY)
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

func (c *Conn) UpdateUser(user *User, newUsername string) error {
	const UPDATE_QUERY = `
	UPDATE users
	SET username = ?
	WHERE id = ? AND username = ?
	`

	ctx := context.Background()

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(UPDATE_QUERY)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newUsername, user.Id, user.Username)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func (c *Conn) UpdatePassword(user *User) error {
	const UPDATE_QUERY = `
	UPDATE users
	SET password = ?
	WHERE id = ? AND username = ?
	`

	ctx := context.Background()

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(UPDATE_QUERY)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Password, user.Id, user.Username)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func (c *Conn) DeleteUser(user *User) error {
	const DELETE_QUERY = `
	DELETE FROM users
	WHERE id = ? AND username = ?
	`

	ctx := context.Background()

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(DELETE_QUERY)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Id, user.Username)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}
