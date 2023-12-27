package service

import (
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/szymon676/codehund/types"
)

type UserServicer interface {
	CreateUser(*types.User) error
}

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(userin *types.User) error {
	err := checkUserStruct(userin)
	if err != nil {
		return err
	}
	user, err := correctUserStruct(userin)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("insert into users (username, email, password) values ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func checkUserStruct(user *types.User) error {
	if len(user.Username) < 3 || len(user.Username) > 16 {
		return ErrInvalidUserNameLength
	}
	if strings.TrimSpace(user.Username) == "" {
		return ErrInvalidUsername
	}
	if strings.TrimSpace(user.Email) == "" {
		return ErrInvalidEmail
	}
	if strings.TrimSpace(user.Password) == "" {
		return ErrInvalidPassword
	}
	if len(user.Email) < 8 || len(user.Email) > 50 {
		return ErrInvalidEmailLength
	}
	if len(user.Password) < 4 || len(user.Password) > 200 {
		return ErrInvalidPasswordLength
	}
	return nil
}

func correctUserStruct(user *types.User) (*types.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(user.Password)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &types.User{
		Username: strings.TrimSpace(user.Username),
		Email:    strings.TrimSpace(user.Email),
		Password: string(hashedPassword),
	}, nil
}

var (
	ErrInvalidUsername       = errors.New("invalid username")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrInvalidUserNameLength = errors.New("invalid username length")
	ErrInvalidEmailLength    = errors.New("invalid email length")
	ErrInvalidPasswordLength = errors.New("invalid password length")
	ErrPasswordsDoNotMatch   = errors.New("passwords do not match")
)
