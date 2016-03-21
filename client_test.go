package outreach

import (
	"testing"
	"time"
)

var (
	clientId     = ""
	clientSecret = ""
	redirectURL  = ""
	accessToken  = "2dd462117dae04761c3051b6059b300390f5caeeb529753f8927f314144a33d1"
	refreshToken = "0023eeb0343e6b282cbbf6839e4611e85af890a31ed0952f9ec54d404a6db02f"
	expires      = "2016-03-10T09:38:55Z"
)

func TestClient(t *testing.T) {

	var conf Config
	conf.ClientId = clientId
	conf.ClientSecret = clientSecret
	conf.RedirectURL = redirectURL
	conf.Scopes = []string{"read_sequences"}

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
	conf.Scopes = []string{"read_sequences"}

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
		t.Errorf("Sequence Error: %s", err)
	}

	if len(seq.Data) == 0 {
		t.Fail()
	}

}
