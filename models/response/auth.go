package response

import "github.com/duyledat197/go-template/models/domain"

// Login ...
type Login struct {
	User    domain.User `json:"user"`
	Token   string      `json:"accessToken"`
	Success bool        `json:"success"`
}

// Register ...
type Register struct {
	Success bool `json:"success"`
}
