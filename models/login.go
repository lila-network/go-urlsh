package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Login struct {
	bun.BaseModel `bun:"table:logins"`
	Cookie        string    `bun:"cookie,pk,type:uuid"`
	UserName      string    `bun:"username,notnull"`
	Expires       time.Time `bun:"expires,notnull"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember string `json:"remember"`
}
