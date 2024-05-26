package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var secretKeyTest = []byte("Tajna")

// CREATING ACCESS TOKEN
func TestCreateAccessToken(t *testing.T) {
	email := "johndoe@example.com"
	tokenString, err := createAccessToken(email)
	if err != nil {
		t.Errorf("Error creating access token: %v", err)
	}

	if tokenString == "" {
		t.Error("Empty token string returned")
	}
}

// VERIFYING ACCESS TOKEN
func TestVerifyAccessToken(t *testing.T) {
	email := "johndoe@example.com"
	tokenString, err := createAccessToken(email)
	if err != nil {
		t.Errorf("Error creating access token: %v", err)
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		t.Errorf("Error verifying access token: %v", err)
	}

	if !token.Valid {
		t.Error("Access token is not valide")
	} else {
		t.Log("Access token valide")
	}
}

// CREATING REFRESH TOKEN
func TestCreateRefreshToken(t *testing.T) {
	email := "johndoe@example.com"
	refreshToken, err := createRefreshToken(email)
	if err != nil {
		t.Errorf("Error creating refresh token: %v", err)
	}

	if refreshToken == "" {
		t.Error("Empty token string returned")
	}
}

// VERIFYING REFERSH TOKEN
func VerifyRefreshToken(t *testing.T) {
	email := "johndoe@example.com"
	refreshTokenString, err := createRefreshToken(email)
	if err != nil {
		t.Errorf("Error creating refresh token: %v", err)
	}

	refreshToken, err := verifyToken(refreshTokenString)
	if err != nil {
		t.Errorf("Error verifying refresh token: %v", err)
	}

	if !refreshToken.Valid {
		t.Error("Refresh token is not valide")
	} else {
		t.Log("Access token valide")
	}
}

// AUTE WHEN ALL COOKIES ARE PRESENT
func TestAllCookiesAndTokens(t *testing.T) {
	// Create a sample request to test middleware
	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

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
	SetCookie(w, "AccessToken", accessToken, expiration)
	SetCookie(w, "RefreshToken", refreshToken, expiration)
	// Retrieve the cookies from the recorder and set them on the request
	cookies := w.Result().Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	handler := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code: %d, got: %d", http.StatusOK, w.Code)
	}

	foundAccessCookie := false
	foundRefreshCookie := false
	foundAccessToken := false
	foundRefreshToken := false
	for _, cookie := range cookies {
		if cookie.Name == "AccessToken" {
			foundAccessCookie = true
		}
		if cookie.Value == accessToken {
			foundAccessToken = true
		}
		if cookie.Name == "RefreshToken" {
			foundRefreshCookie = true
		}
		if cookie.Value == refreshToken {
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
	t.Log("Test successful")
}

// AUTH WHEN ACCESS COOKIE MISSING
func TestWhenAccessCookieMissing(t *testing.T) {

	// Create a sample request to test middleware
	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	email := "johndoe@example.com"
	expiration := time.Now().Add(time.Hour)
	refreshToken, err := createRefreshToken(email)
	if err != nil {
		t.Errorf("Error creating refresh token: %v", err)
	}
	SetCookie(w, "RefreshToken", refreshToken, expiration)
	cookies := w.Result().Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	handler := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Dummy handler do nothing
	}))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code: %d, got: %d", http.StatusOK, w.Code)
	}
	foundRefreshCookie := false
	foundRefreshToken := false
	for _, cookie := range cookies {
		if cookie.Name == "RefreshToken" {
			foundRefreshCookie = true
		}
		if cookie.Value == refreshToken {
			foundRefreshToken = true
		}
	}
	if !foundRefreshCookie {
		t.Error("Refresh Cookie not found")
	}
	if !foundRefreshToken {
		t.Error("Refresh Token not found")
	}
	t.Log("Test successful")
}

// AUTH WHEN REFRESH COOKIE MISSING
func TestWhenRefreshCookieMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	email := "johndoe@example.com"
	expiration := time.Now().Add(time.Hour)
	accessToken, err := createAccessToken(email)
	if err != nil {
		t.Errorf("Error creating access token: %v", err)
	}
	SetCookie(w, "AccessToken", accessToken, expiration)
	cookies := w.Result().Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	handler := authenticateMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Dummy handler do nothing
	}))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code: %d, got: %d", http.StatusOK, w.Code)
	}
	foundAccessCookie := false
	foundAccessToken := false
	for _, cookie := range cookies {
		if cookie.Name == "AccessToken" {
			foundAccessCookie = true
		}
		if cookie.Value == accessToken {
			foundAccessToken = true
		}
	}
	if !foundAccessCookie {
		t.Error("Access Cookie not found")
	}
	if !foundAccessToken {
		t.Error("Access Token not found")
	}
	t.Log("Test successful")
}
