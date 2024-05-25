package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

// CREATE NEW EVENT
func TestCreateNewEvent(t *testing.T) {
	ConnectToDB()
	t.Run("CreateExistingEvent", func(t *testing.T) {
		event := &EventsSchema{
			Date:     time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC),
			Title:    "Example Event",
			FromDate: time.Date(2024, time.May, 20, 10, 0, 0, 0, time.UTC),
			ToDate:   time.Date(2024, time.May, 20, 12, 0, 0, 0, time.UTC),
		}

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		createNewEvent(rr, req)

		if status := rr.Code; status != http.StatusConflict {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusConflict)
		}

		expected := "A event with the same title already exists"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateExistingEvent")

		}
	})

	t.Run("CreateParametarNotAvailable", func(t *testing.T) {
		event := &EventsSchema{
			Date:     time.Time{},
			Title:    "",
			FromDate: time.Time{},
			ToDate:   time.Time{},
		}

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		createNewEvent(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

		expected := "You have to input Date,Title,FromDate,ToDate of event"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateParametarNotAvailable")
		}
	})

	t.Run("CreateNewButton", func(t *testing.T) {

		event := &EventsSchema{
			Date:     time.Date(2024, time.May, 21, 0, 0, 0, 0, time.UTC),
			Title:    "Example Event1",
			FromDate: time.Date(2024, time.May, 21, 10, 0, 0, 0, time.UTC),
			ToDate:   time.Date(2024, time.May, 21, 12, 0, 0, 0, time.UTC),
		}

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		createNewEvent(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"message":"Event item created successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateNewButton")
		}
	})
}

// RETURN ALL EVENTS
func TestReturningAllEvents(t *testing.T) {
	ConnectToDB()

	event := []*EventsSchema{
		{
			Date:     time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC),
			Title:    "Example Event",
			FromDate: time.Date(2024, time.May, 20, 10, 0, 0, 0, time.UTC),
			ToDate:   time.Date(2024, time.May, 20, 12, 0, 0, 0, time.UTC),
		},
		{
			Date:     time.Date(2024, time.May, 21, 0, 0, 0, 0, time.UTC),
			Title:    "Example Event1",
			FromDate: time.Date(2024, time.May, 21, 10, 0, 0, 0, time.UTC),
			ToDate:   time.Date(2024, time.May, 21, 12, 0, 0, 0, time.UTC),
		},
	}

	req, err := http.NewRequest("GET", "/event", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	returnAllEvents(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var receivedBody []*EventsSchema

	if err := json.Unmarshal(rr.Body.Bytes(), &receivedBody); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Compare the response body with the expected body
	if !compereEventsReturns(receivedBody, event) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			getEventsAsString(receivedBody), getEventsAsString(event))
	} else {
		t.Logf("%s: Test passed successfully", "TestReturningAllButtons")
	}

}

// RETURNING ONE EVENT
func TestReturnOneEvent(t *testing.T) {
	ConnectToDB()

	id := "6644ecc90e67c9ea133e004c"

	event := &EventsSchema{
		Date:     time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC),
		Title:    "Example Event",
		FromDate: time.Date(2024, time.May, 20, 10, 0, 0, 0, time.UTC),
		ToDate:   time.Date(2024, time.May, 20, 12, 0, 0, 0, time.UTC),
	}

	// Setup
	req, err := http.NewRequest("GET", "/event/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set a valid ID parameter in the request URL
	req = mux.SetURLVars(req, map[string]string{"id": id})

	rr := httptest.NewRecorder()

	// Invoke the handler
	returnOneEvent(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Parse the response body into a single TodoSchema
	var receivedBody EventsSchema
	if err := json.Unmarshal(rr.Body.Bytes(), &receivedBody); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Compare the response body with the expected body
	if !compareEvents(&receivedBody, event) { // Pass pointer to receivedBody
		t.Errorf("Handler returned unexpected body: got %v want %v",
			receivedBody, *event) // Dereference example
	} else {
		t.Logf("%s: Test passed successfully", "ReturnOneButton")
	}
}

// DELETEING ONE EVENT
func TestDeleteOneEvent(t *testing.T) {
	ConnectToDB()

	id := "6644ed6dd6a71aaf01a1a7c0"
	// Setup
	req, err := http.NewRequest("DELETE", "/event/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": id})
	rr := httptest.NewRecorder()

	// Invoke the handler
	deleteOneEvent(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	example := `{"message":"Event item deleted successfully"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != example {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			actual, example)
	} else {
		t.Logf("%s: Test passed successfully", "DeleteOneButton")
	}
}

// UPDATE ONE EVENT
func TestUpdateOneEvent(t *testing.T) {
	ConnectToDB()

	t.Run("UpdateParamatarsNotAvailable", func(t *testing.T) {
		id := "6644ecc90e67c9ea133e004c"

		// Prepare the request body
		event := &EventsSchema{
			Date:     time.Time{},
			Title:    "",
			FromDate: time.Time{},
			ToDate:   time.Time{},
		}

		// Convert todo to JSON
		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}
		// Create the HTTP request
		req, err := http.NewRequest("PATCH", "/event/"+id, bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Set URL parameters
		req = mux.SetURLVars(req, map[string]string{"id": id})

		// Create a recorder to record the response
		rr := httptest.NewRecorder()

		// Invoke the handler
		updateOneEvent(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Check the response body
		expected := "You have to input the Date or Title or ToDate or FromDate to update todo"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateParamatarsNotAvailable")
		}
	})
	t.Run("UpdateOneEvent", func(t *testing.T) {

		id := "6644ecc90e67c9ea133e004c"

		// Prepare the request body

		event := &EventsSchema{
			Title:    "UpdateExample",
			Date:     time.Date(2019, time.May, 20, 0, 0, 0, 0, time.UTC),
			FromDate: time.Date(2019, time.May, 20, 10, 0, 0, 0, time.UTC),
			ToDate:   time.Date(2019, time.May, 20, 12, 0, 0, 0, time.UTC),
		}

		// Convert todo to JSON
		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}
		// Setup
		req, err := http.NewRequest("PATCH", "/event/"+id, bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": id})
		rr := httptest.NewRecorder()

		// Invoke the handler
		updateOneEvent(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body
		example := `{"message":"Event item updated successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != example {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, example)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateOneEvent")
		}
	})

}
