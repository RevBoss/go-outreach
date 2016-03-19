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
	si := &SequenceInstance{}

	_, e := si.Get()
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
	si := &SequenceInstance{
		Client: client,
	}

	s, e := si.Get()
	Fail(t, e)

	if !reflect.DeepEqual(s, expected) {
		Fail(t, fmt.Errorf("Expected: %+v\nGot: %+v\n", expected, s))
	}
}
