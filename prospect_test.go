package outreach_test

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/revboss/go-outreach"
	"reflect"
	"testing"
)

func TestProspectWithoutClient(t *testing.T) {
	pi := &ProspectInstance{}

	e := pi.Post(Prospect{})
	if e == nil {
		Fail(t, errors.New("Prospect should error if a client is not provided."))
	}
}

func TestProspectPost(t *testing.T) {
	expected := Prospect{}

	j, _ := json.Marshal(expected)
	pi := &ProspectInstance{
		Client: MockClient(j),
	}

	e := pi.Post(expected)
	Fail(t, e)
}

func TestProspectGet(t *testing.T) {
	expected := Prospect{
		Data: ProspectData{
			Attributes: ProspectAttributes{
				Address: ProspectAddress{
					Street: []string{"123 Fake St"},
				},
			},
		},
	}

	j, _ := json.Marshal(expected)
	pi := &ProspectInstance{
		Client: MockClient(j),
	}

	p, e := pi.Get(0)
	Fail(t, e)

	if !reflect.DeepEqual(p, expected) {
		Fail(t, fmt.Errorf("Got: %+v\nExpected: %+v\n", p, expected))
	}
}
