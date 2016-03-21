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
	oc := &OutreachClient{}

	e := oc.PostProspect(Prospect{})
	if e == nil {
		Fail(t, errors.New("Prospect should error if a client is not provided."))
	}
}

func TestProspectPost(t *testing.T) {
	expected := Prospect{}

	j, _ := json.Marshal(expected)
	oc := &OutreachClient{
		Client: MockClient(j),
	}

	e := oc.PostProspect(expected)
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
	oc := &OutreachClient{
		Client: MockClient(j),
	}

	p, e := oc.GetProspect(0)
	Fail(t, e)

	if !reflect.DeepEqual(p, expected) {
		Fail(t, fmt.Errorf("Got: %+v\nExpected: %+v\n", p, expected))
	}
}
