package request

// Login ...
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register ...
type Register struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
