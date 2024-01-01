package http

import (
	"detaskify/internal/users"
	"log"
	"net/http"
)

// HandleGitHubLogin redirects the user to the GitHub login page
func (h *Handler) HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := h.OAuthService.configs["github"].AuthCodeURL(h.OAuthService.state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGitHubCallback processes the OAuth callback from GitHub
func (h *Handler) HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	// Define the GitHub API URL for user info
	githubAPIURL := "https://api.github.com/user"

	user, err := h.OAuthService.GetUserInfo("github", state, code, githubAPIURL)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	detaskifyUser := users.User{
		Username:     user.Login,
		ProfilePhoto: user.AvatarURL,
		Email:        user.Email,
		Provider:     "GitHub",
		IsVerified:   true,
	}

	err = h.Users.CreateUser(r.Context(), &detaskifyUser)
	// Handle the error from CreateUser
	if err != nil {
		log.Println(err.Error())
		// Redirect or handle the error appropriately
	}
}
