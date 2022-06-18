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

func GetTodos(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	listID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	items, err := dbConn.GetTodoItems(userID, listID)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(items)
}

func GetTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	itemID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	var item models.TodoItem
	item.Id = itemID
	item.UserId = userID

	err = dbConn.GetTodoItem(&item)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(item)
}

func AddTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	item.UserId = userID
	item.Desc = strings.TrimSpace(item.Desc)
	if item.Desc == "" {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	err = dbConn.AddTodoItem(&item)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func UpdateTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	item.UserId = userID
	err = dbConn.UpdateTodoItem(&item)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func DeleteTodo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var item models.TodoItem

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		httpMessage(w, http.StatusBadRequest)
		return
	}

	item.UserId = userID
	err = dbConn.DeleteTodoItem(&item)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetLists(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	lists, err := dbConn.GetTodoLists(userID)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}

func GetList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)
	value := ps.ByName("id")

	listID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	list, err := dbConn.GetTodoList(listID, userID)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func AddList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	list.UserId = userID

	exists, err := dbConn.ListExists(&list)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("list with that name already exists")
		return
	}

	list.UserId = userID
	err = dbConn.AddTodoList(&list)
	if err != nil {
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func UpdateList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	list.UserId = userID
	err = dbConn.UpdateTodoList(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func DeleteList(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userID := req.Context().Value(idKey).(int64)

	var list models.TodoList

	err := json.NewDecoder(req.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusBadRequest)
		return
	}

	list.UserId = userID
	err = dbConn.DeleteTodoList(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	err = dbConn.DeleteTodoItemsFromList(&list)
	if err != nil {
		log.Println(err)
		httpMessage(w, http.StatusInternalServerError)
		return
	}

	httpMessage(w, http.StatusOK)
}
