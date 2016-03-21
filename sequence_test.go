package outreach_test

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/revboss/go-outreach"
	"reflect"
	"testing"
)

func TestSequenceGetWithoutClient(t *testing.T) {
	oc := &OutreachClient{}

	_, e := oc.GetSequences()
	if e == nil {
		Fail(t, errors.New("Sequence should error if a client is not provided."))
	}
}

func TestSequenceGet(t *testing.T) {
	expected := SequenceResponse{
		Data:  []SequenceData{},
		Meta:  SequenceMeta{},
		Links: SequenceLinks{},
	}

	j, _ := json.Marshal(expected)
	client := MockClient(j)
	oc := &OutreachClient{
		Client: client,
	}

	s, e := oc.GetSequences()
	Fail(t, e)

	if !reflect.DeepEqual(s, expected) {
		Fail(t, fmt.Errorf("Expected: %+v\nGot: %+v\n", expected, s))
	}
}
