package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Login struct {
	bun.BaseModel `bun:"table:logins"`
	ID            int       `bun:"id,pk,autoincrement"`
	UserName      string    `bun:"username,notnull"`
	Cookie        string    `bun:"cookie,notnull"`
	Expires       time.Time `bun:"expires,notnull"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
