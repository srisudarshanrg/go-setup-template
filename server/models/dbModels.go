package models

import "time"

// User is a model for a user object in a database, if the application includes authentication
type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	JoinDate  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
