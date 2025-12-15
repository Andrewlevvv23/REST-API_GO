package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"r_d/models"
	"r_d/repository"
	"strconv"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, http.StatusOK, models.ErrorResponse{
		Success: true,
		Message: "Server is healthy",
	})
}

func (h *UserHandler) User(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Use GET")
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		sendErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	sendJSONResponse(w, http.StatusOK, models.GetUserResponse{
		Success: true,
		Message: "User retrieved successfully",
		User:    []models.User{*user},
	})
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Use GET")
		return
	}

	users, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	sendJSONResponse(w, http.StatusOK, models.GetUsersResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Count:   len(users),
		Users:   users,
	})
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Use POST")
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format: "+err.Error())
		return
	}
	defer r.Body.Close()

	if user.Name == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}
	if user.Age <= 14 || user.Age > 100 {
		sendErrorResponse(w, http.StatusBadRequest, "Age must be between 14 and 100")
		return
	}
	if user.Phone == "" || len(user.Phone) < 10 || len(user.Phone) > 15 {
		sendErrorResponse(w, http.StatusBadRequest, "Phone is required :(")
		return
	}

	userID, err := h.repo.Create(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	sendJSONResponse(w, http.StatusCreated, models.CreateUserResponse{
		Success: true,
		Message: "User created successfully",
		UserID:  userID,
	})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		sendErrorResponse(w, http.StatusBadRequest, "Method not allowed. Use PUT")
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format: "+err.Error())
		return
	}

	if user.ID == 0 {
		sendErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}

	UserID, err := h.repo.Update(user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	sendJSONResponse(w, http.StatusOK, models.UpdateUserResponse{
		Success: true,
		Message: "User updated successfully",
		UserID:  UserID,
	})

	//log.Printf("DEBUG: id from request = %+v\n", user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendErrorResponse(w, http.StatusBadRequest, "Method not allowed. Use DELETE")
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format: "+err.Error())
		return
	}

	if user.ID == 0 {
		sendErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}

	err = h.repo.Delete(user.ID)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	sendJSONResponse(w, http.StatusOK, models.ErrorResponse{
		Success: true,
		Message: fmt.Sprintf("User deleted successfully. ID:  %d", user.ID),
	})
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	sendJSONResponse(w, statusCode, models.ErrorResponse{
		Success: false,
		Message: message,
	})
}
