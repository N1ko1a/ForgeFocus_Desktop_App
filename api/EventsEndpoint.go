package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"time"
)

var eventCollection string = "events"
var eventTestCollection string = "events_test"

// Get all events
func returnAllEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllEvents")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authenticateMiddleware := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Collection, ctx, err := connectToCollection(eventCollection)
		if err != nil {
			http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
			return
		}

		cursor, err := Collection.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, "Error finding events", http.StatusInternalServerError)
			return
		}

		var events []*EventsSchema
		if err = cursor.All(ctx, &events); err != nil {
			http.Error(w, "Error retriving elements from cursor", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(events)
	}))

	// Izvršavanje middleware-a
	authenticateMiddleware.ServeHTTP(w, r)

}

// Get one event
func returnOneEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnOneEvent")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authenticateMiddleware := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Collection, ctx, err := connectToCollection(eventCollection)
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

		var event *EventsSchema
		err = Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&event)
		if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(event)
	}))

	// Izvršavanje middleware-a
	authenticateMiddleware.ServeHTTP(w, r)
}

// Create one event
func createNewEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewEvent")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	authenticateMiddleware := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Collection, ctx, err := connectToCollection(eventCollection)
		if err != nil {
			http.Error(w, "Error connecting to collection", http.StatusInternalServerError)
			return
		}

		var event *EventsSchema
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// Unmarshal the request body into the todo variable
		err = json.Unmarshal(reqBody, &event)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		if event.Date == (time.Time{}) && event.Title == "" && event.FromDate == (time.Time{}) && event.ToDate == (time.Time{}) {
			http.Error(w, "You have to input Date,Title,FromDate,ToDate of event", http.StatusBadRequest)
			return
		}
		// Check if a todo with the same content already exists
		existingTodo := Collection.FindOne(ctx, bson.M{"title": event.Title})
		if existingTodo.Err() == nil {
			http.Error(w, "A event with the same title already exists", http.StatusConflict)
			return
		}

		// If no existing todo is found with the same content, proceed to create a new one
		_, err = Collection.InsertOne(ctx, event)
		if err != nil {
			http.Error(w, "Error inserting new event", http.StatusInternalServerError)
			return
		}

		// Send a JSON response indicating success
		response := map[string]string{"message": "Event item created successfully"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error creating response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}))

	// Izvršavanje middleware-a
	authenticateMiddleware.ServeHTTP(w, r)
}

// Delete one event
func deleteOneEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteOneEvent")
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authenticateMiddleware := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Collection, ctx, err := connectToCollection(eventCollection)
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
			http.Error(w, "Error deleting event", http.StatusInternalServerError)
			return
		}

		response := map[string]string{"message": "Event item deleted successfully"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error creating response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}))

	// Izvršavanje middleware-a
	authenticateMiddleware.ServeHTTP(w, r)
}

// Update one event
func updateOneEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateOneEvent")
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authenticateMiddleware := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Collection, ctx, err := connectToCollection(eventCollection)
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

		var event *EventsSchema
		err = json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusInternalServerError)
			return
		}

		if event.Date == (time.Time{}) && event.Title == "" && event.ToDate == (time.Time{}) && event.FromDate == (time.Time{}) {
			http.Error(w, "You have to input the Date or Title or ToDate or FromDate to update todo", http.StatusBadRequest)
			return
		}
		updateFields := bson.M{}
		if !event.Date.IsZero() {
			updateFields["date"] = event.Date
		}
		if event.Title != "" {
			updateFields["title"] = event.Title
		}
		if !event.FromDate.IsZero() {
			updateFields["fromdate"] = event.FromDate
		}
		if !event.ToDate.IsZero() {
			updateFields["todate"] = event.ToDate
		}

		filter := bson.M{"_id": objID}
		update := bson.M{"$set": updateFields}
		_, err = Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			http.Error(w, "Error updating event", http.StatusInternalServerError)
			return
		}

		response := map[string]string{"message": "Event item updated successfully"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error creating response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}))

	// Izvršavanje middleware-a
	authenticateMiddleware.ServeHTTP(w, r)
}
