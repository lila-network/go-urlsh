package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Link struct {
	bun.BaseModel `bun:"table:links"`
	ID            string    `bun:"id,pk" json:"id,omitempty"`
	Created       time.Time `bun:"created" json:"created,omitempty"`
	Modified      time.Time `bun:"modified" json:"modified,omitempty"`
	URL           string    `bun:"url,notnull" json:"url"`
	Description   string    `bun:"description" json:"description,omitempty"`
}
