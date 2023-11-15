package data

type User struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
