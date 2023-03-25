package repositories

import "loginRegisterTemplate/types"

func CreateUser(user types.User) error {
	return Connection().Create(&user).Error
}

func GetUser(user types.User) (types.User, error) {
	var foundUser types.User
	err := Connection().Where(&user).First(&foundUser).Error
	if err != nil {
		return types.User{}, err
	}
	return foundUser, nil
}
