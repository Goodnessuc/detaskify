package http

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

var (
	WakaTimeOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     "xSWXLNNeZdBRRjkeZBbMoPIS",
		ClientSecret: "waka_sec_lId6UhJAUqb0mWL2g9BjHMLP3CQxvBHzixi8JY1GK642HY9Z8yJFMR4XqFtkTlFieIHspv8GbmqYcOzH",
		Scopes:       []string{"email", "read_stats"}, // Define required scopes
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://wakatime.com/oauth/authorize",
			TokenURL: "https://wakatime.com/oauth/token",
		},
	}
	oauthStateString = os.Getenv("") // Replace with a random state string for production
)

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
