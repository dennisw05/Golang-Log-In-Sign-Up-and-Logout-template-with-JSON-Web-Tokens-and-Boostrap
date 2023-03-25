package helpers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" && r.URL.Path != "/signup" && !strings.HasPrefix(r.URL.Path, "/public/") {
			tokenString, err := GetToken(r)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			claims, err := ValidateToken(tokenString)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			username, ok := claims["username"].(string)
			if !ok {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			ctx := context.WithValue(r.Context(), "username", username)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func GenerateToken(username string, expires time.Time) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expires.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func GetToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func SetJWTToken(username string, w http.ResponseWriter) error {
	expires := time.Now().Add(time.Hour * 24)
	token, err := GenerateToken(username, expires)
	if err != nil {
		return fmt.Errorf("unexpected error setting jwt")
	}
	cookie := http.Cookie{
		Name:     "jwt-token",
		Value:    token,
		Expires:  expires,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}
