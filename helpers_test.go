package outreach_test

import (
	"encoding/json"
	"errors"
	. "github.com/revboss/go-outreach"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

type Body []byte
type Transport struct {
	body Body
}
type AddProspectTransport Transport

func Fail(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
		t.Fail()
	}
}

func MockClient(body []byte) *http.Client {
	hc := &http.Client{
		Transport: Transport{body},
	}

	return hc
}

func MockAddProspectClient(body []byte) *http.Client {
	hc := &http.Client{
		Transport: AddProspectTransport{body},
	}

	return hc
}

func (b Body) Read(t []byte) (n int, err error) {
	if len(t) < len(b) {
		return len(t), nil
	}

	copy(t, b)
	return len(b), io.EOF
}

func (b Body) Close() error {
	return nil
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		Body: t.body,
	}

	if req.Body != nil {
		_, e := ioutil.ReadAll(req.Body)
		if e != nil {
			return resp, e
		}
	}

	return resp, nil
}

func (ap AddProspectTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body == nil {
		return nil, errors.New("Expected request body")
	}

	_, e := ioutil.ReadAll(req.Body)
	if e != nil {
		return nil, e
	}

	sr := SequenceAddProspectResponse{}

	e = json.Unmarshal(ap.body, &sr)
	if e != nil {
		return nil, e
	}

	split := strings.Split(req.URL.Path, "/")
	id, e := strconv.Atoi(split[len(split)-1])
	if e != nil {
		return nil, errors.New("ID must be an integer")
	}
	sr.Data.ID = id

	j, e := json.Marshal(sr)
	if e != nil {
		return nil, e
	}

	return &http.Response{
		Body: Body(j),
	}, nil
}
