package http

import (
	"detaskify/internal/users"
	"log"
	"net/http"
)

// HandleGitHubLogin redirects the user to the GitHub login page
func (h *Handler) HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := GitHubOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGitHubCallback processes the OAuth callback from GitHub
func (h *Handler) HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserInfo(r.URL.Query().Get("state"), r.URL.Query().Get("code"), "")
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	detaskifyUser := users.Users{
		Username:     user.Login,
		ProfilePhoto: user.AvatarURL,
		Email:        user.Email,
		Provider:     "GitHub",
		IsVerified:   true,
	}

	err = h.Users.CreateUser(r.Context(), &detaskifyUser)
}
