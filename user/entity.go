package user

import "time"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Avatar     string
	Phone      string
	Role       Role
	created_at time.Time
	updated_at time.Time
}
