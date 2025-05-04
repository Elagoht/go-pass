package models

import "time"

// Account model for the database
type Account struct {
	Id         int       `json:"id"`
	Platform   string    `json:"platform" validate:"required,min=1,max=100"`
	URL        string    `json:"url" validate:"required,url,max=255"`
	Identity   string    `json:"identity" validate:"required,min=1,max=255"`
	Passphrase string    `json:"passphrase" validate:"required,min=8"`
	Notes      string    `json:"notes" validate:"max=1000"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
