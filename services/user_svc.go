package services

import (
	"go-api/domains"
	"go-api/dtos/userdtos"
	"go-api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: ur,
	}
}

func (us *UserService) CreateUser(req userdtos.CreateUserRequest) (*userdtos.CreateUserResponse, error) {
	// Hash password from request
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Map requestBody to user model
	user := domains.User{
		Email:        req.Email,
		PasswordHash: string(hashed[:]),
	}

	// Insert DB row
	err = us.UserRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	// Map result to responsebody
	responseBody := userdtos.CreateUserResponse{
		ID: user.ID.String(),
	}

	return &responseBody, err
}
