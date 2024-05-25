package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

var buttonsCollection string = "buttons"
var buttonsTestCollection string = "buttons_test"

// Get all buttons
func returnAllButtons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllButtons")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Collection, ctx, err := connectToCollection(buttonsTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
		return
	}

	cursor, err := Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error finding buttons", http.StatusInternalServerError)
		return
	}

	var buttons []*ButtonsSchema
	if err = cursor.All(ctx, &buttons); err != nil {
		http.Error(w, "Error retriving elements from cursor", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(buttons)

}

// Get one
func returnOneButton(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnOneButton")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalide ID format", http.StatusBadRequest)
		return
	}

	Collection, ctx, err := connectToCollection(buttonsTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to Collection", http.StatusInternalServerError)
		return
	}

	var button *ButtonsSchema

	err = Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&button)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(button)
}

// Create button
func createNewButton(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewButton")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Collection, ctx, err := connectToCollection(buttonsTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to Collection", http.StatusInternalServerError)
		return
	}

	var button *ButtonsSchema
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(reqBody, &button)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		return
	}

	if button.Name == "" {
		http.Error(w, "You have to input the name of a button", http.StatusBadRequest)
		return
	}

	existingButton := Collection.FindOne(ctx, bson.M{"name": button.Name})
	if existingButton.Err() == nil {
		http.Error(w, "A button with the same name already exists", http.StatusConflict)
		return
	}

	_, err = Collection.InsertOne(ctx, button)
	if err != nil {
		http.Error(w, "Error inserting new todo", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Button item created successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creatign response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Delete one button
func deleteOneButton(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteOneButton")
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Collection, ctx, err := connectToCollection(buttonsTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to Collection", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalide ID format", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": objID}
	_, err = Collection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, "Error deleting button", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Button item deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Update button
func updateOneButton(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateOneButton")
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Collection, ctx, err := connectToCollection(buttonsTestCollection)
	if err != nil {
		http.Error(w, "Error connecting to Collection", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalide ID format", http.StatusBadRequest)
		return
	}

	var button *ButtonsSchema
	err = json.NewDecoder(r.Body).Decode(&button)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if button.Name == "" {
		http.Error(w, "You have to input name to update the button", http.StatusBadRequest)
		return
	}

	update := bson.M{}
	if button.Name != "" {
		update["$set"] = bson.M{"name": button.Name}
	}

	filter := bson.M{"_id": objID}
	_, err = Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error updating todo", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Button item updated successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
