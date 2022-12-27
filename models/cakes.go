package models

import (
	"database/sql"
	"time"
)

type Cakes struct {
	ID          int          `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	Rating      int          `json:"rating" db:"rating"`
	Image       string       `json:"image" db:"image"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
