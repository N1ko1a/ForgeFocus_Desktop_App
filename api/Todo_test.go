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

// CREATE NEW TODO
func TestCreateNewTodo(t *testing.T) {
	ConnectToDB()

	//This will create new todo if it already dosent exist!!!!
	t.Run("CreateExistingTodo", func(t *testing.T) {

		todo := &TodoSchema{
			Content:   "Test Todo",
			Workspace: "Test Workspace",
		}

		// Convert todo to JSON
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatal(err)
		}

		// Setup
		req, err := http.NewRequest("POST", "/todo", bytes.NewBuffer(todoJSON))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		email := "johndoe@example.com"
		expiration := time.Now().Add(time.Hour)
		accessToken, err := createAccessToken(email)
		if err != nil {
			t.Errorf("Error creating access token: %v", err)
		}
		refreshToken, err := createRefreshToken(email)
		if err != nil {
			t.Errorf("Error creating refresh token: %v", err)
		}
		SetCookie(rr, "AccessToken", accessToken, expiration)
		SetCookie(rr, "RefreshToken", refreshToken, expiration)
		// Retrieve the cookies from the recorder and set them on the request
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		// Invoke the handler
		createNewTodo(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusConflict {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}

		// Check the response body
		expected := "A todo with the same content already exists"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateExistingTodo")
		}
	})
	t.Run("CreateOneTodo", func(t *testing.T) {
		// Create a *TodoSchema instance
		todo := &TodoSchema{
			Content:   "Test Todo1",
			Workspace: "Test Workspace",
		}

		// Convert todo to JSON
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatal(err)
		}

		// Setup
		req, err := http.NewRequest("POST", "/todo", bytes.NewBuffer(todoJSON))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		email := "johndoe@example.com"
		expiration := time.Now().Add(time.Hour)
		accessToken, err := createAccessToken(email)
		if err != nil {
			t.Errorf("Error creating access token: %v", err)
		}
		refreshToken, err := createRefreshToken(email)
		if err != nil {
			t.Errorf("Error creating refresh token: %v", err)
		}
		SetCookie(rr, "AccessToken", accessToken, expiration)
		SetCookie(rr, "RefreshToken", refreshToken, expiration)
		// Retrieve the cookies from the recorder and set them on the request
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		// Invoke the handler
		createNewTodo(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body
		expected := `{"message":"Todo item created successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateOneTodo")
		}
	})
}

// RETURN ALL TODOS
func TestReturningAllTodos(t *testing.T) {
	ConnectToDB()

	examples := []*TodoSchema{
		{
			Content:   "Test Todo",
			Workspace: " UpdateWorksapce",
			Completed: false,
		},
		{
			Content:   "UpdateTodo",
			Workspace: "UpdateWorkspace",
			Completed: false,
		},
	}

	// Setup
	req, err := http.NewRequest("GET", "/todo", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	email := "johndoe@example.com"
	expiration := time.Now().Add(time.Hour)
	accessToken, err := createAccessToken(email)
	if err != nil {
		t.Errorf("Error creating access token: %v", err)
	}
	refreshToken, err := createRefreshToken(email)
	if err != nil {
		t.Errorf("Error creating refresh token: %v", err)
	}
	SetCookie(rr, "AccessToken", accessToken, expiration)
	SetCookie(rr, "RefreshToken", refreshToken, expiration)
	// Retrieve the cookies from the recorder and set them on the request
	cookies := rr.Result().Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	// Invoke the handler
	returnAllTodos(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Parse the response body into a slice of maps
	var receivedBody []*TodoSchema
	if err := json.Unmarshal(rr.Body.Bytes(), &receivedBody); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}
	// Compare the response body with the expected body
	if !compereTodosReturns(receivedBody, examples) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			getTodosAsString(receivedBody), getTodosAsString(examples))
	} else {
		t.Logf("%s: Test passed successfully", "TestReturningAllTodos")
	}
}

// RETURNING ONE TODO
func TestReturningOneTodo(t *testing.T) {
	ConnectToDB()

	id := "6644a596586b29b22517ebab"

	example := &TodoSchema{
		Content:   "Test Todo1",
		Workspace: "Test Workspace",
		Completed: false,
	}

	// Setup
	req, err := http.NewRequest("GET", "/todo/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set a valid ID parameter in the request URL
	req = mux.SetURLVars(req, map[string]string{"id": id})

	rr := httptest.NewRecorder()

	email := "johndoe@example.com"
	expiration := time.Now().Add(time.Hour)
	accessToken, err := createAccessToken(email)
	if err != nil {
		t.Errorf("Error creating access token: %v", err)
	}
	refreshToken, err := createRefreshToken(email)
	if err != nil {
		t.Errorf("Error creating refresh token: %v", err)
	}
	SetCookie(rr, "AccessToken", accessToken, expiration)
	SetCookie(rr, "RefreshToken", refreshToken, expiration)
	// Retrieve the cookies from the recorder and set them on the request
	cookies := rr.Result().Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	// Invoke the handler
	returnOneTodo(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Parse the response body into a single TodoSchema
	var receivedBody TodoSchema // Not a pointer
	if err := json.Unmarshal(rr.Body.Bytes(), &receivedBody); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Compare the response body with the expected body
	if !compareTodos(&receivedBody, example) { // Pass pointer to receivedBody
		t.Errorf("Handler returned unexpected body: got %v want %v",
			receivedBody, *example) // Dereference example
	} else {
		t.Logf("%s: Test passed successfully", "Success")
	}
}

// DELETE ONE TODO
func TestDeleteOneTodo(t *testing.T) {
	ConnectToDB()

	id := "6644a596586b29b22517ebab"
	// Setup
	req, err := http.NewRequest("DELETE", "/todo/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": id})
	rr := httptest.NewRecorder()

	email := "johndoe@example.com"
	expiration := time.Now().Add(time.Hour)
	accessToken, err := createAccessToken(email)
	if err != nil {
		t.Errorf("Error creating access token: %v", err)
	}
	refreshToken, err := createRefreshToken(email)
	if err != nil {
		t.Errorf("Error creating refresh token: %v", err)
	}
	SetCookie(rr, "AccessToken", accessToken, expiration)
	SetCookie(rr, "RefreshToken", refreshToken, expiration)
	// Retrieve the cookies from the recorder and set them on the request
	cookies := rr.Result().Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	// Invoke the handler
	deleteTodo(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	example := `{"message":"Todo item deleted successfully"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != example {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			actual, example)
	} else {
		t.Logf("%s: Test passed successfully", "DeleteOneTodo")
	}
}

// DELETE ALL TODOS
func TestDeleteAllTodos(t *testing.T) {
	ConnectToDB()

	workspace := "Test Workspace"
	// Setup
	req, err := http.NewRequest("DELETE", "/todo/"+workspace, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"workspace": workspace})
	rr := httptest.NewRecorder()

	email := "johndoe@example.com"
	expiration := time.Now().Add(time.Hour)
	accessToken, err := createAccessToken(email)
	if err != nil {
		t.Errorf("Error creating access token: %v", err)
	}
	refreshToken, err := createRefreshToken(email)
	if err != nil {
		t.Errorf("Error creating refresh token: %v", err)
	}
	SetCookie(rr, "AccessToken", accessToken, expiration)
	SetCookie(rr, "RefreshToken", refreshToken, expiration)
	// Retrieve the cookies from the recorder and set them on the request
	cookies := rr.Result().Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	// Invoke the handler
	deleteAllTodos(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	example := `{"message":"Todo items deleted successfully"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != example {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			actual, example)
	} else {
		t.Logf("%s: Test passed successfully", "DeleteAllTodos")
	}
}

// TEST UPDATE TODO
func TestUpdateOneTodo(t *testing.T) {
	ConnectToDB()

	t.Run("UpdateParamatarsNotAvailable", func(t *testing.T) {
		id := "664500d16daeeceacd944324"

		// Prepare the request body
		todo := &TodoSchema{
			Content:   "",
			Workspace: "",
		}

		// Convert todo to JSON
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatal(err)
		}
		// Create the HTTP request
		req, err := http.NewRequest("PATCH", "/todo/"+id, bytes.NewBuffer(todoJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Set URL parameters
		req = mux.SetURLVars(req, map[string]string{"id": id})

		// Create a recorder to record the response
		rr := httptest.NewRecorder()

		email := "johndoe@example.com"
		expiration := time.Now().Add(time.Hour)
		accessToken, err := createAccessToken(email)
		if err != nil {
			t.Errorf("Error creating access token: %v", err)
		}
		refreshToken, err := createRefreshToken(email)
		if err != nil {
			t.Errorf("Error creating refresh token: %v", err)
		}
		SetCookie(rr, "AccessToken", accessToken, expiration)
		SetCookie(rr, "RefreshToken", refreshToken, expiration)
		// Retrieve the cookies from the recorder and set them on the request
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		// Invoke the handler
		updateTodo(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Check the response body
		expected := "You have to input the workspace or content to update todo"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateParamatarsNotAvailable")
		}
	})
	t.Run("UpdateOneTodo", func(t *testing.T) {

		id := "664500d16daeeceacd944324"

		// Prepare the request body

		todo := &TodoSchema{
			Content:   "UpdateTodo",
			Workspace: "UpdateWorkspace",
		}

		// Convert todo to JSON
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatal(err)
		}
		// Setup
		req, err := http.NewRequest("PATCH", "/todo/"+id, bytes.NewBuffer(todoJSON))
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": id})
		rr := httptest.NewRecorder()

		email := "johndoe@example.com"
		expiration := time.Now().Add(time.Hour)
		accessToken, err := createAccessToken(email)
		if err != nil {
			t.Errorf("Error creating access token: %v", err)
		}
		refreshToken, err := createRefreshToken(email)
		if err != nil {
			t.Errorf("Error creating refresh token: %v", err)
		}
		SetCookie(rr, "AccessToken", accessToken, expiration)
		SetCookie(rr, "RefreshToken", refreshToken, expiration)
		// Retrieve the cookies from the recorder and set them on the request
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		// Invoke the handler
		updateTodo(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body
		example := `{"message":"Todo item updated successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != example {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, example)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateOneTodo")
		}
	})

}

// UPDATE ALL TODOS
func TestUpdateAllTodos(t *testing.T) {
	ConnectToDB()

	t.Run("UpdateWorkspaceNotAvailable", func(t *testing.T) {

		workspace := ""

		// Prepare the request body
		todo := &TodoSchema{
			Workspace: "",
		}

		// Convert todo to JSON
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatal(err)
		}
		// Create the HTTP request
		req, err := http.NewRequest("PATCH", "/todo/"+workspace, bytes.NewBuffer(todoJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Set URL parameters
		req = mux.SetURLVars(req, map[string]string{"workspace": workspace})

		// Create a recorder to record the response
		rr := httptest.NewRecorder()

		email := "johndoe@example.com"
		expiration := time.Now().Add(time.Hour)
		accessToken, err := createAccessToken(email)
		if err != nil {
			t.Errorf("Error creating access token: %v", err)
		}
		refreshToken, err := createRefreshToken(email)
		if err != nil {
			t.Errorf("Error creating refresh token: %v", err)
		}
		SetCookie(rr, "AccessToken", accessToken, expiration)
		SetCookie(rr, "RefreshToken", refreshToken, expiration)
		// Retrieve the cookies from the recorder and set them on the request
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		// Invoke the handler
		updateAllTodos(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Check the response body
		expected := "You have to input the name of workspace witch you wont  to update todo"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateParamatarsNotAvailable")
		}
	})

	t.Run("UpdateNewWorkspaceNotAvailable", func(t *testing.T) {

		workspace := "Test Workspace"

		// Prepare the request body
		todo := &TodoSchema{
			Workspace: "",
		}

		// Convert todo to JSON
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatal(err)
		}

		// Create the HTTP request
		req, err := http.NewRequest("PATCH", "/todo/"+workspace, bytes.NewBuffer(todoJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Set URL parameters
		req = mux.SetURLVars(req, map[string]string{"workspace": workspace})

		// Create a recorder to record the response
		rr := httptest.NewRecorder()

		email := "johndoe@example.com"
		expiration := time.Now().Add(time.Hour)
		accessToken, err := createAccessToken(email)
		if err != nil {
			t.Errorf("Error creating access token: %v", err)
		}
		refreshToken, err := createRefreshToken(email)
		if err != nil {
			t.Errorf("Error creating refresh token: %v", err)
		}
		SetCookie(rr, "AccessToken", accessToken, expiration)
		SetCookie(rr, "RefreshToken", refreshToken, expiration)
		// Retrieve the cookies from the recorder and set them on the request
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		// Invoke the handler
		updateAllTodos(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Check the response body
		expected := "You have to input the new name of workspace"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateParamatarsNotAvailable")
		}
	})

	t.Run("UpdateAllTodos", func(t *testing.T) {

		workspace := "Test Workspace"

		// Prepare the request body
		todo := &TodoSchema{
			Workspace: " UpdateWorksapce",
		}

		// Convert todo to JSON
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatal(err)
		}

		// Setup
		req, err := http.NewRequest("PATCH", "/todo/"+workspace, bytes.NewBuffer(todoJSON))
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"workspace": workspace})
		rr := httptest.NewRecorder()

		email := "johndoe@example.com"
		expiration := time.Now().Add(time.Hour)
		accessToken, err := createAccessToken(email)
		if err != nil {
			t.Errorf("Error creating access token: %v", err)
		}
		refreshToken, err := createRefreshToken(email)
		if err != nil {
			t.Errorf("Error creating refresh token: %v", err)
		}
		SetCookie(rr, "AccessToken", accessToken, expiration)
		SetCookie(rr, "RefreshToken", refreshToken, expiration)
		// Retrieve the cookies from the recorder and set them on the request
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		// Invoke the handler
		updateAllTodos(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body
		example := `{"message":"Todo items updated successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != example {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, example)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateOneTodo")
		}
	})

}
