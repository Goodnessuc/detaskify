package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
