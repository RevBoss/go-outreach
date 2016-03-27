package outreach

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SequenceResponse struct {
	Data   []SequenceData
	Meta   SequenceMeta
	Links  SequenceLinks
	Errors []map[string]interface{}
}

type SequenceData struct {
	ID         int
	Type       string
	Attributes SequenceAttributes
}

type SequenceAttributes struct {
	Name    string
	Created string
	Updated string
}

type SequenceMeta struct {
	Page    SequencePage
	Results SequenceResults
}

type SequencePage struct {
	Current int
	Entries int
	Maximum int
}

type SequenceResults struct {
	Total int
}

type SequenceLinks struct {
	Next string
	Prev string
	Self string
}

type SequenceAddProspect struct {
	Data SequenceAddProspectData `json:"data"`
}

type SequenceAddProspectData struct {
	Relationships SequenceAddProspectRelationships `json:"relationships"`
}

type SequenceAddProspectRelationships struct {
	Prospects []SequenceAddProspectProspect `json:"prospects"`
}

type SequenceAddProspectProspect struct {
	Data SequenceAddProspectProspectData `json:"data"`
}

type SequenceAddProspectProspectData struct {
	ID string `json:"id"`
}

type SequenceAddProspectResponse struct {
	Data   SequenceData
	Links  SequenceLinks
	Errors []map[string]interface{}
}

type SequenceInstance struct {
	Client *http.Client
}

func (s *SequenceInstance) AddProspect(id int, pids ...int) (SequenceAddProspectResponse, error) {
	seq := SequenceAddProspectResponse{}

	if s.Client == nil {
		return seq, errors.New("You must assign a HTTP client.")
	}

	sp := SequenceAddProspect{}

	for _, v := range pids {
		sp.Data.Relationships.Prospects =
			append(sp.Data.Relationships.Prospects, SequenceAddProspectProspect{
				Data: SequenceAddProspectProspectData{
					ID: strconv.Itoa(v),
				},
			})
	}

	j, e := json.Marshal(sp)
	if e != nil {
		return seq, e
	}

	req, e := http.NewRequest("PATCH", "https://api.outreach.io/1.0/sequences/"+strconv.Itoa(id), bytes.NewReader(j))
	if e != nil {
		return seq, e
	}

	resp, e := s.Client.Do(req)
	if e != nil {
		return seq, e
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return seq, e
	}

	e = json.Unmarshal(body, &seq)
	if e != nil {
		return seq, e
	}

	if len(seq.Errors) > 0 {
		return seq, fmt.Errorf("Got error response: %+v\n", seq.Errors)
	}

	return seq, nil
}

func (s *SequenceInstance) Get(opts ...int) (SequenceResponse, error) {
	seq := SequenceResponse{}

	if s.Client == nil {
		return seq, errors.New("You must assign a HTTP client.")
	}

	resp, e := s.Client.Get("https://api.outreach.io/1.0/sequences")
	if e != nil {
		return seq, e
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return seq, e
	}

	e = json.Unmarshal(body, &seq)
	if e != nil {
		return seq, e
	}

	if len(seq.Errors) > 0 {
		return seq, fmt.Errorf("Got error response: %+v\n", seq.Errors)
	}

	return seq, nil
}
