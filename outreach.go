package outreach

import (
	"errors"
	"net/http"
	"time"
)

type OutreachClient struct {
	Client *http.Client
}

func Configure(clientId string, clientSecret string, redirectURL string, scopes []string) (Config, error) {

	var conf Config

	if len(clientId) == 0 || len(clientSecret) == 0 || len(redirectURL) == 0 || len(scopes) == 0 {
		err := errors.New("Missing required configuration parameters")
		return conf, err
	}

	conf.ClientId = clientId
	conf.ClientSecret = clientSecret
	conf.RedirectURL = redirectURL
	conf.Scopes = scopes

	return conf, nil
}

func (c Config) NewOutreachClient(accessToken string, refreshToken string, tokenExpires string) (*OutreachClient, error) {

	oc := &OutreachClient{}

	if len(accessToken) == 0 || len(refreshToken) == 0 || len(tokenExpires) == 0 {
		err := errors.New("Missing required client parameters")
		return nil, err
	}

	var creds Credentials
	creds.AccessToken = accessToken
	creds.RefreshToken = refreshToken

	texp, err := time.Parse(time.RFC3339, tokenExpires)
	if err != nil {
		err := errors.New("Invalid tokenExpires time format - requires RFC3339")
		return nil, err
	}
	creds.TokenExpires = texp

	client, err := Client(c, creds)
	if err != nil {
		return nil, err
	}

	oc.Client = client
	return oc, nil
}
