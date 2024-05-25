package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var userCollection string = "users"

type User struct {
	FirstName      string
	LastName       string
	Email          string
	Password       string
	PasswordConf   string
	AlwaysLoggedIn bool
}

// CREATE USER
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateUser")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	Collection, ctx, err := connectToCollection(userCollection)
	if err != nil {
		http.Error(w, "Error connecting to Collection", http.StatusInternalServerError)
		return
	}

	var user *User
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if user.FirstName == "" && user.LastName == "" && user.Email == "" && user.Password == "" && user.PasswordConf == "" {
		http.Error(w, "You have to input all fields", http.StatusBadRequest)
		return
	}

	if user.FirstName == "" {
		http.Error(w, "You have to input First Name", http.StatusBadRequest)
		return
	}
	if user.LastName == "" {
		http.Error(w, "You have to input Last Name", http.StatusBadRequest)
		return
	}
	if user.Email == "" {
		http.Error(w, "You have to input Email", http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		http.Error(w, "You have to input Password", http.StatusBadRequest)
		return
	}
	if user.PasswordConf == "" {
		http.Error(w, "You have to input Password Confirmation", http.StatusBadRequest)
		return
	}

	existingUser := Collection.FindOne(ctx, bson.M{"email": user.Email})
	if existingUser.Err() == nil {
		http.Error(w, "A user with the same email already exists", http.StatusConflict)
		return
	}

	isValidEmail := govalidator.IsEmail(user.Email)
	if !isValidEmail {
		http.Error(w, "Invalide Email", http.StatusBadRequest)
		return
	}

	isStrongPassword := isStrongPassword(user.Password)
	if !isStrongPassword {
		http.Error(w, "Password not strong enough", http.StatusBadRequest)
		return
	}
	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error while hashing password", http.StatusInternalServerError)
		return
	}
	// Comparing the hashed password with the input password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.PasswordConf))
	if err != nil {
		http.Error(w, "Passwords do not mach ", http.StatusInternalServerError)
		return
	}
	var addUser = bson.M{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"password":  hashedPassword,
	}

	_, err = Collection.InsertOne(ctx, addUser)
	if err != nil {
		http.Error(w, "Error inserting new user", http.StatusInternalServerError)
		return
	}

	tokenString, err := createAccessToken(user.Email)
	if err != nil {
		http.Error(w, "No email found", http.StatusInternalServerError)
		return
	}

	refreshTokenString, err := createRefreshToken(user.Email)
	if err != nil {
		http.Error(w, "No email found", http.StatusInternalServerError)
		return
	}
	SetCookie(w, "AccessToken", tokenString, time.Now().Add(time.Minute*20))
	SetCookie(w, "RefreshToken", refreshTokenString, time.Now().Add(time.Hour*25))

	// Send a JSON response indicating success
	response := map[string]string{"message": "User created successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: login")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Collection, ctx, err := connectToCollection(userCollection)
	if err != nil {
		http.Error(w, "Error connecting to Collection", http.StatusInternalServerError)
		return
	}

	var user User
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "You have to input all fields", http.StatusBadRequest)
		return
	}

	var existingUser UserSchema
	err = Collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Comparing the hashed password with the input password
	err = bcrypt.CompareHashAndPassword(existingUser.Password, []byte(user.Password))
	if err != nil {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	tokenString, err := createAccessToken(user.Email)
	if err != nil {
		http.Error(w, "No email found", http.StatusInternalServerError)
		return
	}

	refreshTokenString, err := createRefreshToken(user.Email)
	if err != nil {
		http.Error(w, "No email found", http.StatusInternalServerError)
		return
	}
	SetCookie(w, "AccessToken", tokenString, time.Now().Add(time.Minute*20))
	SetCookie(w, "RefreshToken", refreshTokenString, time.Now().Add(time.Hour*25))

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the data to JSON format
	responseData := map[string]interface{}{
		"user":  existingUser,
		"token": tokenString,
	}

	// Write the JSON data to the response writer
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

// LOGOUT
func Logout(w http.ResponseWriter, r *http.Request) {
	ClearCookie(w, "jwt")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
