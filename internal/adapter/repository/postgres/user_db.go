// primary adapter for user repository using PostgreSQL
package postgres

import (
	"database/sql"

	"github.com/nocson47/go-hex-concept/internal/core/domain"
	"github.com/nocson47/go-hex-concept/internal/core/port"
)

type UsersRepositoryDB struct {
	db *sql.DB
}

// Update return type to match interface
func NewUsersRepositoryDB(db *sql.DB) port.UserRepository {
	return &UsersRepositoryDB{db: db}
}

func (r *UsersRepositoryDB) Create(user *domain.User) error {
	query := `
        INSERT INTO users (user_name, email, password, phone_number, firstname, lastname, dob, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP)
        RETURNING user_id, created_at`
	return r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.Phone,
		user.FirstName,
		user.LastName,
		user.DoB,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *UsersRepositoryDB) GetAllUsers() ([]*domain.User, error) {
	query := `
		SELECT user_id, user_name, email, password, phone_number, firstname, lastname, dob, created_at
		FROM users
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Phone,
			&user.FirstName,
			&user.LastName,
			&user.DoB,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UsersRepositoryDB) GetByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	query := `
        SELECT user_id, user_name, email, password, phone_number, firstname, lastname, dob, created_at
        FROM users WHERE user_id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.DoB,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UsersRepositoryDB) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `
        SELECT user_id, user_name, email, password, phone_number, firstname, lastname, dob, created_at,
        FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.DoB,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UsersRepositoryDB) Update(user *domain.User) error {
	query := `
        UPDATE users 
        SET user_name = $1, email = $2, phone_number = $3, firstname = $4, lastname = $5, dob = $6, updated_at = CURRENT_TIMESTAMP
        WHERE user_id = $7`
	result, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.Phone,
		user.FirstName,
		user.LastName,
		user.DoB,
		user.ID,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *UsersRepositoryDB) Delete(id int64) error {
	query := `DELETE FROM users WHERE user_id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *UsersRepositoryDB) List() ([]*domain.User, error) {
	query := `
        SELECT user_id, user_name, email, password, phone_number, firstname, lastname, dob, created_at
        FROM users ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Phone,
			&user.FirstName,
			&user.LastName,
			&user.DoB,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
