package service

import (
	"errors"
	"go-blog-management-system/model"
	"go-blog-management-system/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// Register a new user
func (service *UserService) RegisterUser(user *model.User) (*model.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("invalid user data")
	}

	// Check if the email is already registered
	existingUser, _ := service.UserRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	// Create the user
	createdUser, err := service.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// Authenticate a user (login)
func (service *UserService) AuthenticateUser(email, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := service.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

// Fetch user by ID
func (service *UserService) GetUser(id int) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	user, err := service.UserRepo.GetUser(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// Fetch all users
func (service *UserService) GetAllUsers() ([]model.User, error) {
	users, err := service.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

// Update user details
func (service *UserService) UpdateUser(user *model.User) (*model.User, error) {
	if user.ID <= 0 || user.Name == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("invalid user data")
	}

	existingUser, err := service.UserRepo.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	updatedUser, err := service.UserRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// Delete a user by ID
func (service *UserService) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	existingUser, err := service.UserRepo.GetUser(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return service.UserRepo.DeleteUser(id)
}
