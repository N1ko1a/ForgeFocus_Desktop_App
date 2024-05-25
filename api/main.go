package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupTodoRoutes(router *mux.Router) {
	todoRouter := router.PathPrefix("/todo").Subrouter()
	todoRouter.HandleFunc("", createNewTodo).Methods("POST")
	todoRouter.HandleFunc("/{id}", returnOneTodo).Methods("GET")
	todoRouter.HandleFunc("", returnAllTodos).Methods("GET")
	todoRouter.HandleFunc("/{id}", deleteTodo).Methods("DELETE")
	todoRouter.HandleFunc("/{id}", updateTodo).Methods("PATCH")
	todoRouter.HandleFunc("/{workspace}", deleteAllTodos).Methods("DELETE")
	todoRouter.HandleFunc("/{workspace}", updateAllTodos).Methods("PATCH")
}

func setupButtonRoutes(router *mux.Router) {
	buttonRouter := router.PathPrefix("/button").Subrouter()
	buttonRouter.HandleFunc("", returnAllButtons).Methods("GET")
	buttonRouter.HandleFunc("/{id}", returnOneButton).Methods("GET")
	buttonRouter.HandleFunc("", createNewButton).Methods("POST")
	buttonRouter.HandleFunc("/{id}", deleteOneButton).Methods("DELETE")
	buttonRouter.HandleFunc("/{id}", updateOneButton).Methods("PATCH")
}

func setupEventRoutes(router *mux.Router) {
	buttonRouter := router.PathPrefix("/event").Subrouter()
	buttonRouter.HandleFunc("", returnAllEvents).Methods("GET")
	buttonRouter.HandleFunc("/{id}", returnOneEvent).Methods("GET")
	buttonRouter.HandleFunc("", createNewEvent).Methods("POST")
	buttonRouter.HandleFunc("/{id}", deleteOneEvent).Methods("DELETE")
	buttonRouter.HandleFunc("/{id}", updateOneEvent).Methods("PATCH")
}

func setupUserRoutes(router *mux.Router) {
	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("POST")
}

func main() {
	ConnectToDB()
	myRouter := mux.NewRouter().StrictSlash(true)
	setupTodoRoutes(myRouter)
	setupButtonRoutes(myRouter)
	setupEventRoutes(myRouter)
	setupUserRoutes(myRouter)

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}
