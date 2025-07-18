// users model represents the user entity in the domain
package domain

import "time"

// User represents the user entity in the domain
type User struct {
	ID        int64     `db:"user_id"`
	Username  string    `db:"user_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Phone     *string   `db:"phone_number"`
	FirstName string    `db:"firstname"`
	LastName  string    `db:"lastname"`
	DoB       *string   `db:"dob"`
	CreatedAt time.Time `db:"created_at"`
	// UpdatedAt *time.Time `db:"updated_at"`
}
