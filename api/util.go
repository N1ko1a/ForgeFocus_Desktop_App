package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("your-secret-key")

func isStrongPassword(password string) bool {
	// Provera dužine šifre
	if len(password) < 8 {
		return false
	}

	// Provera prisustva velikih slova, malih slova, brojeva i specijalnih karaktera
	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range password {
		switch {
		//Prolazi kroz celu abecedu i proverava da li je bar jedan karakter veliko slovo
		case 'A' <= char && char <= 'Z':
			hasUpperCase = true
			//Prolazi kroz celu abecedu i proverava da li je bar jedan karakter malo slovo
		case 'a' <= char && char <= 'z':
			hasLowerCase = true
			//Proverava da li se broj nalazi u sifri
		case '0' <= char && char <= '9':
			hasDigit = true
			//Proverava da li ima neki od ovih elemenata
		case regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(string(char)):
			hasSpecialChar = true
		}
	}

	// Provera zadovoljenja svih kriterijuma
	return hasUpperCase && hasLowerCase && hasDigit && hasSpecialChar
}

func createAccessToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,                                   // Subject (user identifier)
		"iss": "ForgeFocus-app",                        // Issuer
		"exp": time.Now().Add(time.Minute * 15).Unix(), // Expiration time
		"iat": time.Now().Unix(),                       // Issued at
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println("New Access Token created")
	return tokenString, nil
}

func createRefreshToken(email string) (string, error) {

	refClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,                                 // Subject (user identifier)
		"iss": "ForgeFocus-app",                      // Issuer
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiration time
		"iat": time.Now().Unix(),                     // Issued at
	})

	refreshTokenString, err := refClaims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println("New Refresh Token created")
	return refreshTokenString, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func SetCookie(w http.ResponseWriter, name string, value string, expiration time.Time) {
	cookie := buildCookie(name, value, expiration)
	fmt.Println("New cookie created")
	http.SetCookie(w, cookie)
}

func ClearCookie(w http.ResponseWriter, name string) {
	// Postavljanje kolačića s istekom
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Hour), // Postavljanje vremena isteka na jedan sat u prošlosti
		Domain:   "localhost",
		Secure:   false,
	}

	// Postavljanje kolačića s istekom na klijentsku stranu
	http.SetCookie(w, cookie)
}

func buildCookie(name string, value string, expires time.Time) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Expires:  expires,
		Domain:   "localhost",
		Secure:   false,
	}

	return cookie
}

func authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var refreshTokenString string
		refershCoockie, err := r.Cookie("RefreshToken")
		if err == http.ErrNoCookie {
			http.Error(w, "Refresh cookie not present", http.StatusUnauthorized)
			return

		} else if err != http.ErrNoCookie {
			fmt.Println("Refresh Coockie present")
			refreshTokenString = refershCoockie.Value

		} else {
			http.Error(w, "Error occurred: "+err.Error(), http.StatusInternalServerError)
			return
		}

		accessCoockie, err := r.Cookie("AccessToken")
		if err == http.ErrNoCookie {
			fmt.Println("Access Coockie not present")
			refreshToken, err := verifyToken(refreshTokenString)
			if err != nil {
				http.Error(w, "Refresh token inside access coockie not present: "+err.Error(), http.StatusUnauthorized)
				return
			}
			refershClaims, ok := refreshToken.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Refresh token claims are invalide", http.StatusUnauthorized)
				return
			}
			email := refershClaims["sub"].(string)

			newAccessToken, err := createAccessToken(email)
			if err != nil {
				http.Error(w, "Error creating new token: "+err.Error(), http.StatusInternalServerError)
				return
			}

			SetCookie(w, "AccessToken", newAccessToken, time.Now().Add(time.Minute*2))

			_, err = jwt.Parse(newAccessToken, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})
			if err != nil {
				http.Error(w, "Error parsing new token: "+err.Error(), http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)

		} else if err != http.ErrNoCookie {
			fmt.Println("Access Coockie present")
			accessTokenString := accessCoockie.Value
			_, err := verifyToken(accessTokenString)
			if err != nil {
				if ve, ok := err.(*jwt.ValidationError); ok {
					if ve.Errors&jwt.ValidationErrorExpired != 0 {
						http.Error(w, "Access token has expired", http.StatusUnauthorized)

						refreshToken, err := verifyToken(refreshTokenString)
						if err != nil {
							http.Error(w, "Refresh token inside access coockie present: "+err.Error(), http.StatusInternalServerError)
							return
						}
						refreshClames, ok := refreshToken.Claims.(jwt.MapClaims)
						if !ok {
							http.Error(w, "refresh token claims are invalide", http.StatusInternalServerError)
							return
						}
						email := refreshClames["sub"].(string)
						newAccessToken, err := createAccessToken(email)
						if err != nil {
							http.Error(w, "Error creating new token: "+err.Error(), http.StatusInternalServerError)
							return
						}
						SetCookie(w, "AccessToken", newAccessToken, time.Now().Add(time.Minute*2))
						_, err = jwt.Parse(newAccessToken, func(token *jwt.Token) (interface{}, error) {
							return secretKey, nil
						})
						if err != nil {
							http.Error(w, "Error parsing new token: "+err.Error(), http.StatusInternalServerError)
							return
						}
						next.ServeHTTP(w, r)
					}
				} else {
					http.Error(w, "Access token not expired and access coockie is present error: "+err.Error(), http.StatusInternalServerError)
					return
				}
			}

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

	})
}
