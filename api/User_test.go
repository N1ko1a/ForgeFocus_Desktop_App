package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type user struct {
	FirstName    string
	LastName     string
	Email        string
	Password     string
	PasswordConf string
}

// CREAT NEW USER
func TestCreateNewUser(t *testing.T) {
	ConnectToDB()

	t.Run("CreatingExistingUser", func(t *testing.T) {
		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "ivanovicn98@gmail.com",
			Password:     "Otaukstream1!",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusConflict {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "A user with the same email already exists"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateExistingTodo")
		}

	})

	t.Run("CreateNewUser", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "ivanovicn98888@gmail.com",
			Password:     "Otaukstream1!",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := `{"message":"User created successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "CreateNewUser")
		}
	})
	t.Run("FirstNameMissing", func(t *testing.T) {

		usr := user{
			FirstName:    "",
			LastName:     "Ivanovic",
			Email:        "ivanovicn98888@gmail.com",
			Password:     "Otaukstream1!",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "You have to input First Name"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "FirstNameMissing")
		}
	})

	t.Run("LastNameMissing", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "",
			Email:        "ivanovicn98888@gmail.com",
			Password:     "Otaukstream1!",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "You have to input Last Name"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "LastNameMissing")
		}
	})

	t.Run("EmailMissing", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "",
			Password:     "Otaukstream1!",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "You have to input Email"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "EmailMissing")
		}
	})
	t.Run("PasswordMissing", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "ivanovicnikola@gmail.com",
			Password:     "",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "You have to input Password"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "PasswordMissing")
		}
	})

	t.Run("PasswordConfMissing", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "ivanovicnikola@gmail.com",
			Password:     "Otaukstream1!",
			PasswordConf: "",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "You have to input Password Confirmation"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "PasswordConfMissing")
		}
	})

	t.Run("EmailNotValid", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "ivanovicnikola",
			Password:     "Otaukstream1!",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "Invalide Email"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "EmailNotValid")
		}
	})

	t.Run("PasswordsNotTheSame", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "ivanovicnikola@gmail.com",
			Password:     "Otaukstream1!11",
			PasswordConf: "Otaukstream1!",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "Passwords do not mach"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "PasswordsNotTheSame")
		}
	})

	t.Run("PasswordsNotStrongEnough", func(t *testing.T) {

		usr := user{
			FirstName:    "Nikola",
			LastName:     "Ivanovic",
			Email:        "ivanovicnikola@gmail.com",
			Password:     "nikola",
			PasswordConf: "nikola",
		}
		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Register(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusConflict)
		}
		// Check the response body
		expected := "Password not strong enough"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "PasswordNotStrongEnough")
		}
	})
}
