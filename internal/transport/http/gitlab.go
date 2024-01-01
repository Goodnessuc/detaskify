package http

import (
	"detaskify/internal/users"
	"log"
	"net/http"
)

func (h *Handler) HandleGitLabLogin(w http.ResponseWriter, r *http.Request) {
	url := h.OAuthService.configs["gitlab"].AuthCodeURL(h.OAuthService.state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGitLabCallback processes the OAuth callback from GitLab
func (h *Handler) HandleGitLabCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	gitLabAPIURL := "https://gitlab.com/api/v4/user" // GitLab API URL for user info

	user, err := h.OAuthService.GetUserInfo("gitlab", state, code, gitLabAPIURL)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	detaskifyUser := users.User{
		Username:     user.Login,
		ProfilePhoto: user.AvatarURL,
		Email:        user.Email,
		Provider:     "GitLab",
		IsVerified:   true,
	}

	err = h.Users.CreateUser(r.Context(), &detaskifyUser)
	// Handle the error from CreateUser
	if err != nil {
		log.Println(err.Error())
		// Redirect or handle the error appropriately
	}
}
