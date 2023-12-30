package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleGitLabLogin redirects the user to the GitLab login page
func HandleGitLabLogin(w http.ResponseWriter, r *http.Request) {
	url := GitLabOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGitLabCallback processes the OAuth callback from GitLab
func HandleGitLabCallback(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserInfo(r.URL.Query().Get("state"), "https://gitlab.com/api/v4/user", r.URL.Query().Get("code"))
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

	w.Write(jsonData)
}
