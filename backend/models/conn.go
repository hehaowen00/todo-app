package models

import (
	"database/sql"
)

type Conn struct {
	db *sql.DB
}

func NewConn(db *sql.DB) (*Conn, error) {
	conn := &Conn{
		db,
	}

	err := conn.Migrate(&User{})
	if err != nil {
		return nil, err
	}

	err = conn.Migrate(&TodoList{})
	if err != nil {
		return nil, err
	}

	err = conn.Migrate(&TodoItem{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Conn) Sql() *sql.DB {
	return c.db
}

type migrate interface {
	Migrate(*sql.DB) error
}

func (c *Conn) Migrate(t migrate) error {
	return t.Migrate(c.db)
}
