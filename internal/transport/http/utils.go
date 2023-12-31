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

type OAuthService struct {
	configs map[string]*oauth2.Config
	client  *http.Client
	state   string
}

func NewOAuthService() *OAuthService {
	return &OAuthService{
		configs: map[string]*oauth2.Config{
			"wakatime": {
				RedirectURL:  os.Getenv("OAUTH_CALLBACK_URL"),
				ClientID:     os.Getenv("WAKATIME_CLIENT_ID"),
				ClientSecret: os.Getenv("WAKATIME_CLIENT_SECRET"),
				Scopes:       []string{"read_logged_time"},
				Endpoint: oauth2.Endpoint{
					AuthURL:  "https://wakatime.com/oauth/authorize",
					TokenURL: "https://wakatime.com/oauth/token",
				},
			},
			"github": {
				RedirectURL:  os.Getenv("OAUTH_CALLBACK_URL"),
				ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
				ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
				Scopes:       []string{"user:email"},
				Endpoint:     github.Endpoint,
			},
			"gitlab": {
				RedirectURL:  os.Getenv("OAUTH_CALLBACK_URL"),
				ClientID:     os.Getenv("GITLAB_CLIENT_ID"),
				ClientSecret: os.Getenv("GITLAB_CLIENT_SECRET"),
				Scopes:       []string{"read_user"},
				Endpoint: oauth2.Endpoint{
					AuthURL:  "https://gitlab.com/oauth/authorize",
					TokenURL: "https://gitlab.com/oauth/token",
				},
			},
		},
		client: &http.Client{},
		state:  os.Getenv("OAUTH_STATE_STRING"),
	}
}

func (s *OAuthService) GetToken(provider, state, code string) (*oauth2.Token, error) {
	if state != s.state {
		return nil, fmt.Errorf("invalid oauth state")
	}

	config, ok := s.configs[provider]
	if !ok {
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}

	return config.Exchange(context.Background(), code)
}

func (s *OAuthService) GetUserInfo(provider, state, code, APIURL string) (ProviderUser, error) {
	var user ProviderUser

	token, err := s.GetToken(provider, state, code)
	if err != nil {
		return user, err
	}

	req, err := http.NewRequest("GET", APIURL, nil)
	if err != nil {
		return user, fmt.Errorf("failed creating request: %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := s.client.Do(req)
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

func JSONContentTypeWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}
