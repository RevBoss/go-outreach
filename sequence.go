package outreach

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type SequenceResponse struct {
	Data  []SequenceData
	Meta  SequenceMeta
	Links SequenceLinks
}

type SequenceData struct {
	ID         int
	Type       string
	Attributes SequenceAttributes
}

type SequenceAttributes struct {
	Name string
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

func (c *OutreachClient) GetSequences(opts ...int) (SequenceResponse, error) {
	seq := SequenceResponse{}

	if c.Client == nil {
		return seq, errors.New("You must assign an Outreach Client.")
	}

	resp, e := c.Client.Get("https://api.outreach.io/1.0/sequences")
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

	return seq, nil
}
