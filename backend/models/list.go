package models

import (
	"context"
	"database/sql"
	"strings"
)

type TodoList struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"-"`
	Name   string `json:"name"`
}

func NewTodoList(userId int64, name string) TodoList {
	var id int64 = 0
	return TodoList{
		id,
		userId,
		name,
	}
}

func (list *TodoList) Trim() {
	list.Name = strings.TrimSpace(list.Name)
}

func (list *TodoList) Migrate(db *sql.DB) error {
	const CREATE_TABLE_QUERY = `
	CREATE TABLE IF NOT EXISTS lists (
		id integer PRIMARY KEY,
		user_id integer not null,
		name text not null default null,
		unique (user_id, name),
		foreign key (user_id) REFERENCES users (id)
	)`

	_, err := db.Exec(CREATE_TABLE_QUERY)

	return err
}

func (c *Conn) ListExists(list *TodoList) (bool, error) {
	const EXISTS_QUERY = `
	SELECT name FROM lists
	WHERE user_id = ? AND name = ?
	`

	stmt, err := c.db.Prepare(EXISTS_QUERY)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRow(list.UserId, list.Name)

	var listName string
	err = row.Scan(&listName)

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, err
}

func (c *Conn) AddTodoList(list *TodoList) error {
	const INSERT_QUERY = `
	INSERT INTO lists (user_id, name)
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

	result, err := stmt.Exec(list.UserId, list.Name)
	if err != nil {
		return nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	list.Id = id

	err = tx.Commit()

	return err
}

func (c *Conn) GetTodoLists(id int64) ([]TodoList, error) {
	const GET_QUERY = `
	SELECT id, name
	FROM lists
	WHERE user_id = ?
	ORDER BY name
	`

	stmt, err := c.db.Prepare(GET_QUERY)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lists := make([]TodoList, 0)

	for rows.Next() {
		var list TodoList

		err = rows.Scan(&list.Id, &list.Name)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	return lists, nil
}

func (c *Conn) GetTodoList(id int64, userId int64) (*TodoList, error) {
	const GET_QUERY = `
	SELECT id, name
	FROM lists
	WHERE id = ? AND user_id = ?
	`

	stmt, err := c.db.Prepare(GET_QUERY)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id, userId)
	if err != nil {
		return nil, err
	}

	var list TodoList

	err = row.Scan(&list.Id, &list.Name)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
func (c *Conn) GetTodoItem(item *TodoItem) error {
	const SELECT_QUERY = `
	SELECT list_id, desc, status
	FROM todos
	WHERE id = ? AND user_id = ?
	`

	stmt, err := c.db.Prepare(SELECT_QUERY)
	if err != nil {
		return err
	}

	row := stmt.QueryRow(item.Id, item.UserId)
	if err != nil {
		return err
	}

	err = row.Scan(&item.ListId, &item.Desc, &item.Status)

	return err
}

func (c *Conn) UpdateTodoList(list *TodoList) error {
	ctx := context.Background()

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	const UPDATE_QUERY = `
	UPDATE lists
	SET name = ?
	WHERE id = ? AND user_id = ?
	`

	stmt, err := tx.Prepare(UPDATE_QUERY)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(list.Name, list.Id, list.UserId)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func (c *Conn) DeleteTodoList(list *TodoList) error {
	const DELETE_QUERY = `
	DELETE FROM lists
	WHERE id = ? AND user_id = ? AND name = ?;
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

	_, err = stmt.Exec(list.Id, list.UserId, list.Name)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}
