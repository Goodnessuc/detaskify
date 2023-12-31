package http

import (
	"detaskify/internal/users"
	"log"
	"net/http"
)

// HandleGitLabLogin redirects the user to the GitLab login page
func (h *Handler) HandleGitLabLogin(w http.ResponseWriter, r *http.Request) {
	url := GitLabOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGitLabCallback processes the OAuth callback from GitLab
func (h *Handler) HandleGitLabCallback(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserInfo(r.URL.Query().Get("state"), "https://gitlab.com/api/v4/user", r.URL.Query().Get("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	detaskifyUser := users.Users{
		Username:     user.Login,
		ProfilePhoto: user.AvatarURL,
		Email:        user.Email,
		Provider:     "GitLab",
		IsVerified:   true,
	}

	err = h.Users.CreateUser(r.Context(), &detaskifyUser)
}
