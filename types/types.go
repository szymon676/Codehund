package types

type PostgresConnectionOptions struct {
	User         string
	DatabaseName string
	Port         string
	Password     string
}

type RedisConnectionOptions struct {
	Port     string
	Password string
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
