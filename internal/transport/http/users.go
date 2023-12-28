package http

import (
	"detaskify/internal/utils"
	"net/http"
)

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
