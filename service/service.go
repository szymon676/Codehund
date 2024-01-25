package service

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/szymon676/codehund/types"
	"golang.org/x/crypto/bcrypt"
)

type Servicer interface {
	GetUserByUsername(string) (*types.User, error)
	FollowUser(int, int) error
	UnfollowUser(int, int) error
	GetFollowers(int) ([]string, error)
	CreateUser(*types.User) error
	CreatePost(*types.Post) error
	DeletePost(int) error
	GetPosts() ([]types.Post, error)
}

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetUserByUsername(username string) (*types.User, error) {
	row, err := s.db.Query("SELECT id, username, email, password FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user types.User
	if row.Next() {
		err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}
	user.Password = ""
	if user.Username == "" {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *Service) getUserByID(id int) (*types.User, error) {
	row, err := s.db.Query("SELECT id, username, email, password FROM users WHERE id= $1", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user types.User
	if row.Next() {
		err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}
	user.Password = ""
	if user.Username == "" {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *Service) GetFollowers(user int) ([]string, error) {
	rows, err := s.db.Query("SELECT followee FROM followers WHERE follower = $1", user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []string
	for rows.Next() {
		var followeeID int
		if err := rows.Scan(&followeeID); err != nil {
			return nil, err
		}

		id := followeeID
		if err != nil {
			return nil, err
		}

		user, err := s.getUserByID(id)
		if err != nil {
			return nil, err
		}

		followers = append(followers, user.Username)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func GetFollowees(user int) ([]string, error) {
	return nil, nil
}

func (s *Service) FollowUser(follower int, followee int) error {
	_, err := s.db.Exec("INSERT INTO followers (follower, followee) VALUES ($1, $2)", follower, followee)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UnfollowUser(follower int, followee int) error {
	_, err := s.db.Exec("DELETE FROM followers WHERE follower=$1 AND followee=$2", follower, followee)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreatePost(post *types.Post) error {
	err := checkPostStruct(post)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("insert into posts (author, content, date) values ($1, $2, $3)", post.Author, post.Content, post.Date)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetPosts() ([]types.Post, error) {
	res, err := s.db.Query("select * from posts")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var posts []types.Post
	for res.Next() {
		var post types.Post
		err := res.Scan(&post.ID, &post.Author, &post.Content, &post.Date)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *Service) DeletePost(id int) error {
	res, err := s.db.Exec("delete from posts where id = $1", id)
	if err != nil {
		return err
	}
	rf, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rf == 0 {
		return errors.New("post not found")
	}
	return nil
}

func (s *Service) CreateUser(userin *types.User) error {
	err := checkUserStruct(userin)
	if err != nil {
		return err
	}

	user, err := correctUserStruct(userin)
	if err != nil {
		return err
	}

	_, err = s.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func checkPostStruct(post *types.Post) error {
	if len(strings.TrimSpace(post.Content)) < 3 {
		return errors.New("post content is to short, minimum 3 characters")
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
