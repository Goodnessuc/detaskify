package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleGitHubLogin redirects the user to the GitHub login page
func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := GitHubOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGitHubCallback processes the OAuth callback from GitHub
func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserInfo(r.URL.Query().Get("state"), r.URL.Query().Get("code"), "")
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
