package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/szymon676/codehund/types"
	"golang.org/x/crypto/bcrypt"
)

type SessionManager struct {
	client *redis.Client
	db     *sql.DB
}

type UserSession struct {
	ID       int
	Username string
	Email    string
}

func NewSessionManager(client *redis.Client, db *sql.DB) *SessionManager {
	return &SessionManager{
		client: client,
		db:     db,
	}
}

func (s *SessionManager) GenerateSession(data *types.User) (string, error) {
	sessionId := uuid.NewString()
	jsonData, _ := json.Marshal(data)
	err := s.client.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func (s *SessionManager) Login(email, password string) (string, error) {
	var user types.User
	err := s.db.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	sessionId := uuid.NewString()
	jsonData, _ := json.Marshal(UserSession{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
	err = s.client.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (s *SessionManager) Logout(sessionId string) error {
	return s.client.Del(context.Background(), sessionId).Err()
}

func (s *SessionManager) GetSession(session string) (*UserSession, error) {
	data, err := s.client.Get(context.Background(), session).Result()
	if err != nil {
		return nil, err
	}

	var userSession UserSession
	err = json.Unmarshal([]byte(data), &userSession)
	if err != nil {
		return nil, err
	}

	return &userSession, nil
}
