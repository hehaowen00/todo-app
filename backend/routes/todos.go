package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"host.local/todo-app/backend/models"
)

func getTodos(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	listID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to parse list ID")
		return
	}

	items, err := dbConn.GetTodoItems(userID, listID)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to get todo items")
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(items)
}

func getTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	itemID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to parse list ID")
		return
	}

	var item models.TodoItem
	item.Id = itemID
	item.UserId = userID

	err = dbConn.GetTodoItem(&item)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to get todo item")
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(item)
}

func addTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	item.UserId = userID
	item.Desc = strings.TrimSpace(item.Desc)

	if item.Desc == "" {
		jsonMessage(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	err = dbConn.AddTodoItem(&item)
	if err != nil {
		jsonMessage(w, http.StatusInternalServerError, "Failed to add todo item")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func updateTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	item.UserId = userID

	err = dbConn.UpdateTodoItem(&item)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to update todo item")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func deleteTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	item.UserId = userID

	err = dbConn.DeleteTodoItem(&item)
	if err != nil {
		jsonMessage(w, http.StatusInternalServerError, "Failed to delete todo item")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getLists(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	lists, err := dbConn.GetTodoLists(userID)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to get todo lists")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}

func getList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	listID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to parse list ID")
		return
	}

	list, err := dbConn.GetTodoList(listID, userID)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to get todo list items")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func addList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	list.UserId = userID

	exists, err := dbConn.ListExists(&list)
	if err != nil {
		jsonMessage(w, http.StatusInternalServerError, "Server Error")
		return
	}

	if exists {
		jsonMessage(w, http.StatusBadRequest, "List already exists")
		return
	}

	list.UserId = userID

	err = dbConn.AddTodoList(&list)
	if err != nil {
		jsonMessage(w, http.StatusInternalServerError, "Failed to add todo list")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func updateList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	list.UserId = userID

	err = dbConn.UpdateTodoList(&list)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to update todo list")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func deleteList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	list.UserId = userID

	err = dbConn.DeleteTodoList(&list)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to delete todo list")
		return
	}

	err = dbConn.DeleteTodoItemsFromList(&list)
	if err != nil {
		log.Println(err)
		jsonMessage(w, http.StatusInternalServerError, "Failed to delete todo list items")
		return
	}

	jsonMessage(w, http.StatusOK, "List deleted")
}
