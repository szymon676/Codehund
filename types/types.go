package types

type ConnectionOptions struct {
	User         string
	DatabaseName string
	Port         string
	Password     string
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Userstate struct {
	Loggedin bool
}

type Post struct {
	Title     string
	Content   string
	CreatedBy uint
	Author    string
}
