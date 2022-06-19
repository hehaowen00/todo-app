package models

import (
	"context"
	"database/sql"
	"errors"
	"strings"
)

type TodoItem struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"-"`
	ListId int64  `json:"list_id"`
	Desc   string `json:"desc"`
	Status bool   `json:"status"`
}

func NewTodoItem(userId int64, listId int64, desc string, status bool) TodoItem {
	var id int64 = 0
	return TodoItem{
		id,
		userId,
		listId,
		desc,
		status,
	}
}

func (item *TodoItem) Trim() {
	item.Desc = strings.TrimSpace(item.Desc)
}

func (item *TodoItem) Migrate(db *sql.DB) error {
	const CREATE_TABLE_QUERY = `
	CREATE TABLE IF NOT EXISTS todos (
		id integer PRIMARY KEY,
		user_id integer not null default null,
		list_id integer not null default null,
		desc text not null,
		status bool not null default false,
		foreign key (user_id) REFERENCES users (id),
		foreign key (list_id) REFERENCES lists (id)
	)`

	_, err := db.Exec(CREATE_TABLE_QUERY)

	return err
}

func (c *Conn) AddTodoItem(item *TodoItem) error {
	const INSERT_QUERY = `
	INSERT INTO todos (user_id, list_id, desc, status)
	VALUES (?, ?, ?, ?)
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

	result, err := stmt.Exec(item.UserId, item.ListId, item.Desc, item.Status)
	if err != nil {
		return nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	item.Id = id

	err = tx.Commit()

	return err
}

func (c *Conn) GetTodoItems(userId int64, listId int64) ([]TodoItem, error) {
	const GET_QUERY = `
	SELECT id, desc, status
	FROM todos
	WHERE user_id = ? AND list_id = ?
	`

	stmt, err := c.db.Prepare(GET_QUERY)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId, listId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]TodoItem, 0)

	for rows.Next() {
		var item TodoItem

		err = rows.Scan(&item.Id, &item.Desc, &item.Status)
		if err != nil {
			return nil, err
		}

		item.ListId = listId
		items = append(items, item)
	}

	return items, nil
}

func (c *Conn) UpdateTodoItem(item *TodoItem) error {
	const UPDATE_QUERY = `
	UPDATE todos
	SET desc = ?, status = ?
	WHERE id = ? and user_id = ? and list_id = ?
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

	result, err := stmt.Exec(item.Desc, item.Status, item.Id, item.UserId, item.ListId)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("failed to update todo")
	}

	err = tx.Commit()

	return err
}

func (c *Conn) DeleteTodoItem(item *TodoItem) error {
	const DELETE_QUERY = `
	DELETE FROM todos
	WHERE id = ? AND user_id = ? AND list_id = ?
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

	_, err = stmt.Exec(item.Id, item.UserId, item.ListId)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func (c *Conn) DeleteTodoItemsFromList(list *TodoList) error {
	const DELETE_QUERY = `
	DELETE FROM todos
	WHERE list_id = ? AND user_id = ?
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

	_, err = stmt.Exec(list.Id, list.UserId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
