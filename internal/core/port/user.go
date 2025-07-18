// users port defines the interfaces for user management
package port

import "github.com/nocson47/go-hex-concept/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
	List() ([]*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}

type UserService interface {
	CreateUser(user *domain.User) error
	GetUser(id int64) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int64) error
	ListUsers() ([]*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}
