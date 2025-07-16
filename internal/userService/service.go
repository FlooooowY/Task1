package userService

import (
	"errors"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req UserRequest) (Users, error)
	GetAllUser() ([]Users, error)
	GetUserByID(id string) (Users, error)
	UpdateUser(id string, req UserRequest) (Users, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) validateUserRequest(req UserRequest) error {
	if req.Email == "" || req.Password == "" {
		return errors.New("email and pass not be empty")
	}
	if !strings.Contains(req.Email, "@") && !strings.Contains(req.Email, ".") {
		return errors.New("email invalid format")
	}
	if len(req.Password) < 8 {
		return errors.New("pass must be at 8 characters long")
	}
	if strings.Contains(req.Password, " ") {
		return errors.New("pass must not contain spaces")
	}

	var hasLetter, hasDigit bool
	for _, ch := range req.Password {
		switch {
		case unicode.IsLetter(ch):
			hasLetter = true
		case unicode.IsDigit(ch):
			hasDigit = true
		}
	}

	if !hasLetter || !hasDigit {
		return errors.New("pass must contain both letters and digits")
	}
	return nil
}

func (s *userService) CreateUser(req UserRequest) (Users, error) {
	if err := s.validateUserRequest(req); err != nil {
		return Users{}, err
	}

	user := Users{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return Users{}, err
	}
	return user, nil
}

func (s *userService) GetAllUser() ([]Users, error) {
	return s.repo.GetAllUser()
}

func (s *userService) GetUserByID(id string) (Users, error) {
	if id == "" {
		return Users{}, errors.New("user ID is required")
	}
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id string, req UserRequest) (Users, error) {
	if id == "" {
		return Users{}, errors.New("user ID is required")
	}

	if err := s.validateUserRequest(req); err != nil {
		return Users{}, err
	}

	user := Users{
		ID:       id,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return Users{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("user id is required")
	}
	return s.repo.DeleteUser(id)
}
