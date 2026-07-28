package links

import (
	"time"

	"github.com/uptrace/bun"
)

// Link represents a link entity in the database.
type Link struct {
	bun.BaseModel `bun:"table:links"`
	ID            string    `bun:"id,pk" json:"id,omitempty"`
	Created       time.Time `bun:"created" json:"created,omitempty"`
	Modified      time.Time `bun:"modified" json:"modified,omitempty"`
	URL           string    `bun:"url,notnull" json:"url"`
	Description   string    `bun:"description" json:"description,omitempty"`
}

// IsValid checks if the Link has a non-empty ID and URL, indicating it is valid.
func (l Link) IsValid() bool {
	return l.ID != "" && l.URL != ""
}
