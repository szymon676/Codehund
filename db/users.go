package db

import "errors"

type User struct {
	ID              uint   `json:"id"`
	UserName        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
}

func checkUserStruct(user *User) error {
	if len(user.UserName) < 5 || len(user.UserName) > 16 {
		return ErrInvalidUserNameLength
	}

	if len(user.Email) < 8 || len(user.Email) > 50 {
		return ErrInvalidEmailLength
	}

	if len(user.Password) < 6 || len(user.Password) > 200 {
		return ErrInvalidPasswordLength
	}

	if user.Password != user.ConfirmPassword {
		return ErrPasswordsDoNotMatch
	}

	return nil
}

var (
	ErrInvalidUserNameLength = errors.New("invalid username length")
	ErrInvalidEmailLength    = errors.New("invalid email length")
	ErrInvalidPasswordLength = errors.New("invalid password length")
	ErrPasswordsDoNotMatch   = errors.New("passwords do not match")
)
