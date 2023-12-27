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

type Userstate struct {
	Loggedin bool
	User     *User
}

type Post struct {
	Title     string
	Content   string
	CreatedBy uint
	Author    string
}
