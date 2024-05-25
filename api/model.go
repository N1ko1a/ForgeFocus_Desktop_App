package main

import "time"

type ButtonsSchema struct {
	Name   string `json:"name"`
	UserId int    `json:"userId"`
}

type EventsSchema struct {
	Date     time.Time `json:"date"`
	Title    string    `json:"title"`
	FromDate time.Time `json:"fromDate"`
	ToDate   time.Time `json:"toDate"`
	UserId   int       `json:"userId"`
}

type TodoSchema struct {
	Content   string `json:"content"`
	Workspace string `json:"workspace"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

type UserSchema struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Password       []byte `json:"password"`
	AlwaysLoggedIn bool   `json:"alwaysLoggedIn"`
}
