package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/melardev/GoMuxGormApiCrud/dtos"
	"github.com/melardev/GoMuxGormApiCrud/models"
	"github.com/melardev/GoMuxGormApiCrud/services"
	"net/http"
	"strconv"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.FetchTodos()
	sendAsJson(w, http.StatusOK, dtos.GetTodoListDto(todos))
}

func GetAllPendingTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.FetchPendingTodos()
	sendAsJson(w, http.StatusOK, dtos.GetTodoListDto(todos))
}
func GetAllCompletedTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.FetchCompletedTodos()
	sendAsJson(w, http.StatusOK, dtos.GetTodoListDto(todos))
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "completed" {
		GetAllCompletedTodos(w, r)
		return
	} else if id == "pending" {
		GetAllPendingTodos(w, r)
		return
	}

	id64, _ := strconv.ParseUint(id, 10, 32)
	todo, err := services.FetchById(uint(id64))
	if err != nil {
		sendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not find Todo"))
		return
	}

	sendAsJson2(w, http.StatusOK, dtos.GetTodoDetaislDto(&todo))
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		sendAsJson(w, http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	defer r.Body.Close()

	todo, err := services.CreateTodo(todo.Title, todo.Description, todo.Completed)
	if err != nil {
		sendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	sendAsJson(w, http.StatusCreated, dtos.GetTodoDetaislDto(&todo))
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must set an ID"))
		return
	}

	var todoInput models.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todoInput); err != nil {
		sendAsJson(w, http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	defer r.Body.Close()

	todo, err := services.UpdateTodo(uint(id), todoInput.Title, todoInput.Description, todoInput.Completed)
	if err != nil {
		sendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	sendAsJson(w, http.StatusOK, dtos.GetTodoDetaislDto(&todo))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must set an ID"))
		return
	}
	todo, err := services.FetchById(uint(id))
	if err != nil {
		sendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("todo not found"))
		return
	}

	err = services.DeleteTodo(&todo)

	if err != nil {
		sendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not delete Todo"))
		return
	}

	sendAsJson(w, http.StatusNoContent, nil)
}

func DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	services.DeleteAllTodos()
	sendAsJson(w, http.StatusNoContent, nil)
}
