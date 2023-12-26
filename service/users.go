package service

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/szymon676/codehund/types"
)

type UserServicer interface {
	CreateUser(*types.User) error
}

type UserService struct {
	db *sqlx.DB
}

func NewUserService(connopts *types.ConnectionOptions) *UserService {
	schema := `
	CREATE TABLE users (
		name text UNIQUE,
		email text UNIQUE,
		password text
	);
	`
	connstring := fmt.Sprintf("port=%s user=%s dbname=%s password=%s sslmode=disable", connopts.Port, connopts.User, connopts.DatabaseName, connopts.Password)
	db, err := sqlx.Connect("postgres", connstring)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("migrated user schemes")
	db.Exec(schema)
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(userin *types.User) error {
	err := checkUserStruct(userin)
	if err != nil {
		return err
	}
	user := correctUserStruct(userin)
	_, err = s.db.Exec("insert into users (name, email, password) values ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func checkUserStruct(user *types.User) error {
	if len(user.Username) < 5 || len(user.Username) > 16 {
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
	if len(user.Password) < 6 || len(user.Password) > 200 {
		return ErrInvalidPasswordLength
	}
	return nil
}

func correctUserStruct(user *types.User) *types.User {
	return &types.User{
		Username: strings.TrimSpace(user.Username),
		Email:    strings.TrimSpace(user.Email),
		Password: strings.TrimSpace(user.Password),
	}
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
