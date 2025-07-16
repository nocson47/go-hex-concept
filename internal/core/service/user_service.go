package service

import (
	"errors"
	"strings"

	"github.com/nocson47/go-hex-concept/internal/core/domain"
	"github.com/nocson47/go-hex-concept/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	// Validate required fields
	if err := s.validateUser(user); err != nil {
		return err
	}

	// Check if email already exists
	existingUser, _ := s.repo.GetByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	return s.repo.Create(user)
}
func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUser(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *domain.User) error {
	if user.ID <= 0 {
		return errors.New("invalid user ID")
	}

	if err := s.validateUser(user); err != nil {
		return err
	}

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int64) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	return s.repo.Delete(id)
}

func (s *UserService) ListUsers() ([]*domain.User, error) {
	return s.repo.List()
}

func (s *UserService) validateUser(user *domain.User) error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}
