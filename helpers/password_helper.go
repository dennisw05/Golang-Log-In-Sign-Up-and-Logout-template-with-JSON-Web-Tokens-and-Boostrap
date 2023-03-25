package helpers

import (
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidPassword(password string) bool {
	re := regexp.MustCompile(`[A-Z]+`)
	if !re.MatchString(password) {
		return false
	}
	re = regexp.MustCompile(`[a-z]+`)
	if !re.MatchString(password) {
		return false
	}
	re = regexp.MustCompile(`[0-9]+`)
	if !re.MatchString(password) {
		return false
	}
	re = regexp.MustCompile(`[^A-Za-z0-9]+`)
	if !re.MatchString(password) {
		return false
	}
	return len(password) >= 6
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswords(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func CheckSignupPassword(password, passwordConfirm string, w http.ResponseWriter) (string, bool) {
	if len(password) < 6 || !IsValidPassword(password) {
		ErrorCookie("signup-error", "Password does not follow the guidelines!", w)
		return "", false
	}
	if password != passwordConfirm {
		ErrorCookie("signup-error", "Password do not match!", w)
		return "", false
	}
	password, err := HashPassword(password)
	if err != nil {
		ErrorCookie("singup-error", "Something went wrong! Please try again.", w)
		return "", false
	}
	return password, true
}
