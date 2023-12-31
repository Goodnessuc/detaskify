package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CodingSession struct {
	Project   string `json:"project"`
	StartTime string `json:"start"`
	EndTime   string `json:"end"`
	Duration  int    `json:"duration"`
}

// FetchUserData fetches user data from WakaTime using the access token
func (h *Handler) FetchUserData(accessToken string) ([]CodingSession, error) {
	apiURL := "https://wakatime.com/api/v1/users/current/durations" // Adjust this URL as needed

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := h.OAuthService.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var sessions []CodingSession
	err = json.Unmarshal(body, &sessions)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return sessions, nil
}

// HandleWakatimeLogin redirects the user to the Wakatime login page
func (h *Handler) HandleWakatimeLogin(w http.ResponseWriter, r *http.Request) {
	url := h.OAuthService.configs["wakatime"].AuthCodeURL(h.OAuthService.state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleWakatimeCallback processes the OAuth callback from Wakatime
func (h *Handler) HandleWakatimeCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")

	token, err := h.OAuthService.GetToken("wakatime", state, code)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	userData, err := h.FetchUserData(token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process the retrieved user data
	// ...
}
