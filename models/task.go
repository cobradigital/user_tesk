package models

import "time"

type Tasks struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
