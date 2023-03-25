package helpers

import (
	"loginRegisterTemplate/repositories"
	"loginRegisterTemplate/types"
	"net/http"

	"gorm.io/gorm"
)

func IsValidUsername(username string) bool {
	user, err := repositories.GetUser(types.User{Username: username})
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	return user.Username != username
}

func CheckSignupUsername(username string, w http.ResponseWriter) bool {
	if len(username) < 5 || len(username) > 20 {
		ErrorCookie("singup-error", "Username is not the correct length!", w)
		return false
	}
	if !IsValidUsername(username) {
		ErrorCookie("signup-error", "Username already in use!", w)
		return false
	}
	return true
}
