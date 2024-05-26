package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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

		//Testing if the cookies and tokens were made
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}

		foundAccessCookie := false
		foundRefreshCookie := false
		foundAccessToken := false
		foundRefreshToken := false
		for _, cookie := range cookies {
			if cookie.Name == "AccessToken" {
				foundAccessCookie = true
			}
			if cookie.Value != "" {
				foundAccessToken = true
			}
			if cookie.Name == "RefreshToken" {
				foundRefreshCookie = true
			}
			if cookie.Value != "" {
				foundRefreshToken = true
			}
		}
		if !foundAccessCookie {
			t.Error("Access Cookie not found")
		}
		if !foundAccessToken {
			t.Error("Access Token not found")
		}
		if !foundRefreshCookie {
			t.Error("Refresh Cookie not found")
		}
		if !foundRefreshToken {
			t.Error("Refresh Token not found")
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

// LOGIN USER
func TestLoginUser(t *testing.T) {
	ConnectToDB()
	t.Run("LoginUser", func(t *testing.T) {
		usr := user{
			Email:    "ivanovicn98@gmail.com",
			Password: "Otaukstream1!",
		}

		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Login(rr, req)

		//Testing if the cookies and tokens were made
		cookies := rr.Result().Cookies()
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}

		foundAccessCookie := false
		foundRefreshCookie := false
		foundAccessToken := false
		foundRefreshToken := false
		for _, cookie := range cookies {
			if cookie.Name == "AccessToken" {
				foundAccessCookie = true
			}
			if cookie.Value != "" {
				foundAccessToken = true
			}
			if cookie.Name == "RefreshToken" {
				foundRefreshCookie = true
			}
			if cookie.Value != "" {
				foundRefreshToken = true
			}
		}
		if !foundAccessCookie {
			t.Error("Access Cookie not found")
		}
		if !foundAccessToken {
			t.Error("Access Token not found")
		}
		if !foundRefreshCookie {
			t.Error("Refresh Cookie not found")
		}
		if !foundRefreshToken {
			t.Error("Refresh Token not found")
		}
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body

		expected := `{"message":"User logeed in successfully"}`
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "LoginUser")
		}
	})

	t.Run("UserNotFound", func(t *testing.T) {
		usr := user{
			Email:    "ivanovicn@gmail.com",
			Password: "Otaukstream1!",
		}

		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Login(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
		// Check the response body

		expected := "User not found"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "UserNotFound")
		}
	})

	t.Run("EmailNotProvided", func(t *testing.T) {
		usr := user{
			Email:    "",
			Password: "Otaukstream1!",
		}

		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Login(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
		// Check the response body

		expected := "You have to input all fields"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "EmailNotProvided")
		}
	})

	t.Run("PasswordNotProvided", func(t *testing.T) {
		usr := user{
			Email:    "ivanovicn98@gmail.com",
			Password: "",
		}

		userJson, err := json.Marshal(usr)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		Login(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
		// Check the response body

		expected := "You have to input all fields"
		actual := strings.TrimSpace(rr.Body.String())
		if actual != expected {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				actual, expected)
		} else {
			t.Logf("%s: Test passed successfully", "EmailNotProvided")
		}
	})
}

// LOGOUT
func TestLogout(t *testing.T) {
	ConnectToDB()
	req, err := http.NewRequest("POST", "/logout", nil)
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

	Logout(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
