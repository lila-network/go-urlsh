package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	UserName      string    `bun:"username,pk" json:"username"`
	Created       time.Time `bun:"created,notnull,default:now()" json:"created"`
	LastLogin     time.Time `bun:"lastlogin" json:"last_login"`
	PasswordSalt  string    `bun:"salt" json:"password_salt"`
	PasswordHash  string    `bun:"password" json:"password_hash"`
}

type UserResponse struct {
	UserName string    `json:"username"`
	Created  time.Time `json:"created"`
}
