package models

import (
	"database/sql"
)

type Conn struct {
	db *sql.DB
}

func NewConn(db *sql.DB) *Conn {
	return &Conn{
		db,
	}
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
