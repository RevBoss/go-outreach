package outreach

import (
	"testing"
	"time"
)

var (
	clientId     = ""
	clientSecret = ""
	redirectURL  = ""
	accessToken  = ""
	refreshToken = ""
	expires      = ""
	scopes       = []string{"CREATE_PROSPECTS_PERMISSION, READ_SEQUENCES_PERMISSION, UPDATE_SEQUENCES_PERMISSION"}
)

func TestClient(t *testing.T) {

	var conf Config
	conf.ClientId = clientId
	conf.ClientSecret = clientSecret
	conf.RedirectURL = redirectURL
	conf.Scopes = scopes

	var creds Credentials
	creds.AccessToken = accessToken
	creds.RefreshToken = refreshToken
	expires := expires

	tt, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		t.Errorf("Expires is empty or invalid: %s, Error: %s", expires, err.Error())
	}
	creds.TokenExpires = tt

	client, err := Client(conf, creds)
	if err != nil || client == nil {
		t.Fail()
	}

}

func TestClientSequences(t *testing.T) {

	var conf Config
	conf.ClientId = clientId
	conf.ClientSecret = clientSecret
	conf.RedirectURL = redirectURL
	conf.Scopes = scopes

	var creds Credentials
	creds.AccessToken = accessToken
	creds.RefreshToken = refreshToken
	expires := expires

	tt, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		t.Errorf("Expires is empty or invalid: %s, Error: %s", expires, err.Error())
	}
	creds.TokenExpires = tt

	client, err := Client(conf, creds)
	if err != nil || client == nil {
		t.Fail()
	}

	si := &SequenceInstance{}
	si.Client = client

	seq, err := si.Get()
	if err != nil {
		t.Errorf("Sequence Error: %w", err)
	}

	if len(seq.Data) == 0 {
		t.Fail()
	}

}
