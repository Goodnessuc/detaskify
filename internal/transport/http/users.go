package http

import (
	"detaskify/internal/users"
	"detaskify/internal/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser users.Users
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.Users.CreateUser(r.Context(), &newUser)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, newUser)
}

// getUserByUsername - Retrieves a user by username
func (h *Handler) getUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	user, err := h.Users.GetUserByUsername(r.Context(), username)
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, user)
}
func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	user, err := h.Users.GetUserByEmail(r.Context(), email)
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusOK, user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	var updateData users.Users
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.Users.UpdateUser(r.Context(), username, &updateData)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, nil)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	err := h.Users.DeleteUser(r.Context(), username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, nil)
}

func (h *Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	// Assuming the username and new password are sent in the request body
	var payload struct {
		Username    string `json:"username"`
		NewPassword string `json:"newPassword"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Validate the input, ensure it's not empty
	if payload.Username == "" || payload.NewPassword == "" {
		utils.ERROR(w, http.StatusBadRequest, errors.New("username and new password are required"))
		return
	}

	// Call the method to update the user's password
	err = h.Users.UpdateUserPassword(r.Context(), payload.Username, payload.NewPassword)
	if err != nil {
		// Handle specific error types if needed
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Password update was successful
	utils.JSON(w, http.StatusOK, map[string]string{"message": "Password updated successfully"})
}
