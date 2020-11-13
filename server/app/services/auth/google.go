package auth

import (
	"context"
	"encoding/json"
	"os"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getClientID() string {
	return os.Getenv("GOOGLE_CLIENT_ID")
}

func getClientSecret() string {
	return os.Getenv("GOOGLE_CLIENT_SECRET")
}

func getRedirectURL() string {
	return os.Getenv("GOOGLE_REDIRECT_URL")
}

func getGoogleOauthConfig(clientID, clientSecret string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
		RedirectURL:  getRedirectURL(),
	}
}

// defaultGoogleOauthConfig todo
func defaultGoogleOauthConfig() *oauth2.Config {
	return getGoogleOauthConfig(
		getClientID(),
		getClientSecret(),
	)
}

// GoogleService todo
type GoogleService interface {
	GenerateState() (string, error)
	GetDefaultAuthCodeURL() (string, string)
	Exchange(code string) (*oauth2.Token, error)
	GetUser(token *oauth2.Token) (*GoogleUserinfoResponse, error)
}

type googleService struct {
	GoogleService
	context      context.Context
	oauth2Config *oauth2.Config
}

// NewGoogleService todo
func NewGoogleService(context context.Context) GoogleService {
	return &googleService{
		context:      context,
		oauth2Config: defaultGoogleOauthConfig(),
	}
}

// GenerateState generate oauth2 state
func (s *googleService) GenerateState() (string, error) {
	str := uuid.NewV4()
	return str.String(), nil
}

// GetDefaultAuthCodeURL todo
func (s *googleService) GetDefaultAuthCodeURL() (string, string) {
	state, err := s.GenerateState()
	if err != nil {
		return "", ""
	}
	url := s.oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url, state
}

// Exchange todo
func (s *googleService) Exchange(code string) (*oauth2.Token, error) {
	token, err := s.oauth2Config.Exchange(s.context, code)
	return token, err
}

// GoogleUserinfoResponse todo
type GoogleUserinfoResponse struct {
	Email         string
	EmailVerified bool
	FamilyName    string
	Gender        string
	GivenName     string
	Hd            string
	Locale        string
	Name          string
	Picture       string
	Profile       string
	Sub           string
}

// GetUser todo
func (s *googleService) GetUser(token *oauth2.Token) (*GoogleUserinfoResponse, error) {
	client := s.oauth2Config.Client(s.context, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	r := &GoogleUserinfoResponse{}
	err = json.NewDecoder(response.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
