package outreach

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type Config struct {
	ClientId     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

type Credentials struct {
	AccessToken  string
	RefreshToken string
	TokenExpires time.Time
}

func Client(conf Config, creds Credentials) (*http.Client, error) {
	oc := &oauth2.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  conf.RedirectURL,
		Scopes:       conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.outreach.io/oauth/authorize",
			TokenURL: "https://api.outreach.io/oauth/token",
		},
	}

	tok := &oauth2.Token{
		AccessToken:  creds.AccessToken,
		RefreshToken: creds.RefreshToken,
		Expiry:       creds.TokenExpires,
	}

	ctx := context.Background()

	client := oc.Client(ctx, tok)

	return client, nil
}
