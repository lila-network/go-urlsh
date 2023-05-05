package models

import (
	"github.com/uptrace/bun"
	"time"
)

type ApiKey struct {
	bun.BaseModel `bun:"table:apikeys"`
	Key           string    `bun:"key,pk,type:uuid,default:gen_random_uuid()" json:"key,omitempty"`
	UserName      string    `bun:"username,notnull" json:"username,omitempty"`
	Created       time.Time `bun:"created,default:now()" json:"created,omitempty"`
	Description   string    `bun:"description,notnull" json:"description"`
}
