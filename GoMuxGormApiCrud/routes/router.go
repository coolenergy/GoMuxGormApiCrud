package routes

import (
	"github.com/gorilla/mux"
	"github.com/melardev/GoMuxGormApiCrud/controllers"
)

var Router *mux.Router
func RegisterRoutes(){
	Router = mux.NewRouter()

	todoApiSubRoute := Router.PathPrefix("/api/todos").Subrouter()

	todoApiSubRoute.HandleFunc("", controllers.GetAllTodos).Methods("GET")
	todoApiSubRoute.HandleFunc("/completed", controllers.GetAllCompletedTodos).Methods("GET")
	todoApiSubRoute.HandleFunc("/pending", controllers.GetAllPendingTodos).Methods("GET")
	todoApiSubRoute.HandleFunc("/{id}", controllers.GetTodoById).Methods("GET")
	todoApiSubRoute.HandleFunc("", controllers.CreateTodo).Methods("POST")
	todoApiSubRoute.HandleFunc("/{id}", controllers.UpdateTodo).Methods("PUT")
	todoApiSubRoute.HandleFunc("/{id}", controllers.UpdateTodo).Methods("PATCH")
	todoApiSubRoute.HandleFunc("/{id}", controllers.DeleteTodo).Methods("DELETE")
	todoApiSubRoute.HandleFunc("", controllers.DeleteAllTodos).Methods("DELETE")
}

