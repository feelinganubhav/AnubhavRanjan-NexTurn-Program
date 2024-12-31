package controller

import (
	"encoding/json"
	"go-blog-management-system/middleware"
	"go-blog-management-system/model"
	"go-blog-management-system/service"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

// Authenticate a user
func (controller *UserController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
		return
	}

	user, err := controller.UserService.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		writeJSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		return
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
		return
	}

	response := map[string]interface{}{
		"message": "Authentication successful!! Save the Bearer Token for CRUD Ops..",
		"user":    user.Name,
		"token":   token,
	}
	writeJSONResponse(w, http.StatusOK, response)
}

// Create a new user
func (controller *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
		return
	}

	_, err := controller.UserService.RegisterUser(&user)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

// Get a user by ID
func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	user, err := controller.UserService.GetUser(userID)
	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	writeJSONResponse(w, http.StatusOK, user)
}

// Get all users
func (controller *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := controller.UserService.GetAllUsers()
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, users)
}

// Update a user
func (controller *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
		return
	}

	user.ID = userID
	updatedUser, err := controller.UserService.UpdateUser(&user)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, updatedUser)
}

// Delete a user by ID
func (controller *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	err = controller.UserService.DeleteUser(userID)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
