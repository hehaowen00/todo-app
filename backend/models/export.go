package models

import "log"

type ListData struct {
	Id    int64      `json:"list_id"`
	Name  string     `json:"name"`
	Items []TodoItem `json:'items"`
}

func (c *Conn) ExportData(userId int64) ([]ListData, error) {
	const SELECT_LISTS_QUERY = `
  SELECT id, name FROM lists
  WHERE user_id = ?
  `

	stmt, err := c.db.Prepare(SELECT_LISTS_QUERY)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	lists := make([]ListData, 0)

	for rows.Next() {
		var list ListData

		err = rows.Scan(&list.Id, &list.Name)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	for idx := range lists {
		items, err := c.GetTodoItems(userId, lists[idx].Id)
		if err != nil {
			return nil, err
		}

		lists[idx].Items = items
	}

	return lists, nil
}
