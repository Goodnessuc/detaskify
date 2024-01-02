package http

import (
	"detaskify/internal/users"
	"detaskify/internal/utils"
	"encoding/json"
	"net/http"
)

// CreateTeam - Handler for creating a new team
func (h *Handler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var newTeam users.Team
	err := json.NewDecoder(r.Body).Decode(&newTeam)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.Teams.CreateTeam(r.Context(), newTeam)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, newTeam)
}

// GetTeamByID - Handler for retrieving a team by ID
func (h *Handler) GetTeamByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	team, err := h.Teams.GetTeamByID(r.Context(), id)
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusOK, team)
}

// UpdateTeam - Handler for updating a team's details
func (h *Handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updateData users.Team
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.Teams.UpdateTeam(r.Context(), id, updateData)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, nil)
}

// DeleteTeam - Handler for deleting a team
func (h *Handler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.Teams.DeleteTeam(r.Context(), id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, nil)
}

// AddUserToTeam - Handler for adding a user to a team
func (h *Handler) AddUserToTeam(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("teamName")
	username := r.URL.Query().Get("username")
	err := h.Teams.AddUserToTeam(r.Context(), teamName, username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, nil)
}

// RemoveUserFromTeam - Handler for removing a user from a team
func (h *Handler) RemoveUserFromTeam(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("teamName")
	username := r.URL.Query().Get("username")
	err := h.Teams.RemoveUserFromTeam(r.Context(), teamName, username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, nil)
}
