package http

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"net/http"
	"os"
)

type ProviderUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

// TODO: CHANGE callback from localhost

var (
	WakaTimeOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     os.Getenv("WAKATIME_CLIENT_ID"),
		ClientSecret: os.Getenv("WAKATIME_CLIENT_SECRET"),
		Scopes:       []string{"email", "read_stats"}, // Define required scopes
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://wakatime.com/oauth/authorize",
			TokenURL: "https://wakatime.com/oauth/token",
		},
	}
	GitHubOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	GitLabOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback", // Update this URL as needed
		ClientID:     "",                                    // Replace with your GitLab Client ID
		ClientSecret: "",                                    // Replace with your GitLab Client Secret
		Scopes:       []string{"read_user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://gitlab.com/oauth/authorize",
			TokenURL: "https://gitlab.com/oauth/token",
		},
	}

	oauthStateString = os.Getenv("") // Replace with a random state string for production
)

func JSONContentTypeWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to JSON
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}

// GetUserInfo retrieves the user information from WakaTime
func GetUserInfo(state string, code, URL string) (ProviderUser, error) {
	var user ProviderUser

	if state != oauthStateString {
		return user, fmt.Errorf("invalid oauth state")
	}

	token, err := WakaTimeOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return user, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return user, fmt.Errorf("failed creating request: %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return user, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	err = json.Unmarshal(contents, &user)
	if err != nil {
		return user, fmt.Errorf("failed unmarshalling user info: %s", err.Error())
	}

	return user, nil
}
