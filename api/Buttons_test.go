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

// Creating new Button
func TestCreateNewButton(t *testing.T) {
	ConnectToDB()
	t.Run("CreateExistingButton", func(t *testing.T) {
		button := &ButtonsSchema{
			Name: "Test button",
		}

		buttonJSON, err := json.Marshal(button)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/button", bytes.NewBuffer(buttonJSON))
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
		createNewButton(rr, req)

		if status := rr.Code; status != http.StatusConflict {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusConflict)
		}

		expected := "A button with the same name already exists"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateExistingTodo")

		}
	})

	t.Run("CreateParametarNotAvailable", func(t *testing.T) {
		button := &ButtonsSchema{
			Name: "",
		}

		buttonJSON, err := json.Marshal(button)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/button", bytes.NewBuffer(buttonJSON))
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
		createNewButton(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

		expected := "You have to input the name of a button"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateParametarNotAvailable")
		}
	})

	t.Run("CreateNewButton", func(t *testing.T) {

		button := &ButtonsSchema{
			Name: "Test Button1",
		}

		buttonJSON, err := json.Marshal(button)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/button", bytes.NewBuffer(buttonJSON))
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
		createNewButton(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"message":"Button item created successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateNewButton")
		}
	})
}

// Returning all buttons
func TestReturningAllButtons(t *testing.T) {
	ConnectToDB()

	example := []*ButtonsSchema{
		{
			Name: "Test Button",
		},
		{
			Name: "Test Button1",
		},
	}

	req, err := http.NewRequest("GET", "/button", nil)
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
	returnAllButtons(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var receivedBody []*ButtonsSchema

	if err := json.Unmarshal(rr.Body.Bytes(), &receivedBody); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Compare the response body with the expected body
	if !compereButtonReturns(receivedBody, example) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			receivedBody, example)
	} else {
		t.Logf("%s: Test passed successfully", "TestReturningAllButtons")
	}

}

// RETURN ONE BUTTON
func TestReturnOneButton(t *testing.T) {
	ConnectToDB()

	id := "6644adbb269e80a7094bab32"

	button := &ButtonsSchema{
		Name: "Test button",
	}

	// Setup
	req, err := http.NewRequest("GET", "/button/"+id, nil)
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
	returnOneButton(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Parse the response body into a single TodoSchema
	var receivedBody ButtonsSchema
	if err := json.Unmarshal(rr.Body.Bytes(), &receivedBody); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Compare the response body with the expected body
	if !compareButtons(&receivedBody, button) { // Pass pointer to receivedBody
		t.Errorf("Handler returned unexpected body: got %v want %v",
			receivedBody, *button) // Dereference example
	} else {
		t.Logf("%s: Test passed successfully", "ReturnOneButton")
	}
}

// DELETEING ONE BUTTON
func TestDeleteOneButton(t *testing.T) {
	ConnectToDB()

	id := "6644adbb269e80a7094bab33"
	// Setup
	req, err := http.NewRequest("DELETE", "/button/"+id, nil)
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
	deleteOneButton(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	example := `{"message":"Button item deleted successfully"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != example {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			actual, example)
	} else {
		t.Logf("%s: Test passed successfully", "DeleteOneButton")
	}
}

// UPDATE ONE BUTTON
func TestUpdateOneButton(t *testing.T) {
	ConnectToDB()

	t.Run("UpdateParamatarsNotAvailable", func(t *testing.T) {
		id := "6644b1878b765cffbb02ceb1"

		// Prepare the request body
		button := &ButtonsSchema{
			Name: "",
		}

		// Convert todo to JSON
		buttonJSON, err := json.Marshal(button)
		if err != nil {
			t.Fatal(err)
		}
		// Create the HTTP request
		req, err := http.NewRequest("PATCH", "/button/"+id, bytes.NewBuffer(buttonJSON))
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
		updateOneButton(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Check the response body
		expected := "You have to input name to update the button"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "UpdateParamatarsNotAvailable")
		}
	})
}
