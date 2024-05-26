package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todoCollection string = "todos"
var todoTestCollection string = "todos_test"

func returnAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTodos")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Provjera autentičnosti korištenjem middleware-a
	authenticateMiddleware := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Collection, ctx, err := connectToCollection(todoTestCollection)
		if err != nil {
			http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
			return
		}

		cursor, err := Collection.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, "Error finding todos", http.StatusInternalServerError)
			return
		}

		var todos []*TodoSchema
		if err = cursor.All(ctx, &todos); err != nil {
			http.Error(w, "Error retrieving elements from cursor", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(todos)
	}))

	// Izvršavanje middleware-a
	authenticateMiddleware.ServeHTTP(w, r)
}

// Get one todo
func returnOneTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Endpoint Hit: returnOneTodo")
	vars := mux.Vars(r)
	id := vars["id"]

	// Convert the id string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	Collection, ctx, err := connectToCollection(todoTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
		return
	}

	var todo *TodoSchema

	// Use the converted ObjectID in the query
	err = Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// Create new todo
func createNewTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Endpoint Hit: createNewTodo")
	Collection, ctx, err := connectToCollection(todoTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
		return
	}

	var todo *TodoSchema
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Unmarshal the request body into the todo variable
	err = json.Unmarshal(reqBody, &todo)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if todo.Content == "" {
		http.Error(w, "You have to input the name of the todo", http.StatusBadRequest)
		return
	}
	if todo.Workspace == "" {
		http.Error(w, "You have to input the workspace that the todo belongs to", http.StatusBadRequest)
		return
	}
	// Check if a todo with the same content already exists
	existingTodo := Collection.FindOne(ctx, bson.M{"content": todo.Content})
	if existingTodo.Err() == nil {
		http.Error(w, "A todo with the same content already exists", http.StatusConflict)
		return
	}

	// If no existing todo is found with the same content, proceed to create a new one
	_, err = Collection.InsertOne(ctx, todo)
	if err != nil {
		http.Error(w, "Error inserting new todo", http.StatusInternalServerError)
		return
	}

	// Send a JSON response indicating success
	response := map[string]string{"message": "Todo item created successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Delete todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Endpoint Hit: deleteTodo")
	Collection, ctx, err := connectToCollection(todoTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	// Convert the id string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": objID}
	_, err = Collection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, "Error deleting todo", http.StatusInternalServerError)
		return
	}

	// Send a JSON response indicating success
	response := map[string]string{"message": "Todo item deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Delete todos
func deleteAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Endpoint Hit: deleteAllTodos")
	Collection, ctx, err := connectToCollection(todoTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	workspace := vars["workspace"]

	if workspace == "" {
		http.Error(w, "You have to input the workspace ", http.StatusBadRequest)
		return
	}

	filter := bson.M{"workspace": workspace}
	_, err = Collection.DeleteMany(ctx, filter)
	if err != nil {
		http.Error(w, "Error deleting todo", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Todo items deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Update todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Endpoint Hit: updateTodo")
	Collection, ctx, err := connectToCollection(todoTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	// Convert the id string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var todo *TodoSchema

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if todo.Content == "" && todo.Workspace == "" {
		http.Error(w, "You have to input the workspace or content to update todo", http.StatusBadRequest)
		return
	}

	updateFields := bson.M{}
	if todo.Content != "" {
		updateFields["content"] = todo.Content
	}
	if todo.Workspace != "" {
		updateFields["workspace"] = todo.Workspace
	}

	update := bson.M{"$set": updateFields}
	filter := bson.M{"_id": objID}
	_, err = Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error updating todo", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Todo item updated successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Update all todo
func updateAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Endpoint Hit: updateAllTodos")
	Collection, ctx, err := connectToCollection(todoTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	workspace := vars["workspace"]

	if workspace == "" {
		http.Error(w, "You have to input the name of workspace witch you wont  to update todo", http.StatusBadRequest)
		return
	}

	var todo *TodoSchema

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	if todo.Workspace == "" {
		http.Error(w, "You have to input the new name of workspace", http.StatusBadRequest)
		return
	}
	update := bson.M{}
	update["$set"] = bson.M{"workspace": todo.Workspace}

	filter := bson.M{"workspace": workspace}
	_, err = Collection.UpdateMany(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error updating todo", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Todo items updated successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
