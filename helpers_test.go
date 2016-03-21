package outreach_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type Body []byte
type Transport struct {
	body Body
}

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
