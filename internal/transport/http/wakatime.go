package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: Add a function to get the user's WakaTime stats

// HandleWakaTimeLogin redirects the user to the WakaTime login page
func HandleWakaTimeLogin(w http.ResponseWriter, r *http.Request) {
	url := WakaTimeOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleWakaTimeCallback processes the OAuth callback from WakaTime
func HandleWakaTimeCallback(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserInfo(r.URL.Query().Get("state"), r.URL.Query().Get("code"), "https://wakatime.com/api/v1/users/current")
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
