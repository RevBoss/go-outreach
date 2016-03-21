package outreach

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Prospect struct {
	Data ProspectData
}

type ProspectData struct {
	Attributes ProspectAttributes
}

type ProspectAttributes struct {
	Address  ProspectAddress
	Company  ProspectCompany
	Contact  ProspectContact
	Personal ProspectPersonal
	Social   ProspectSocial
	Meta     ProspectMeta `json: "metadata"`
}

type ProspectAddress struct {
	City    string
	State   string
	Country string
	Street  []string
	Zip     int
}

type ProspectCompany struct {
	Name     string
	Type     string
	Industry string
	Size     int
	Locality string
}

type ProspectContact struct {
	Timezone string
	Email    string
	Phone    ProspectPhone
}

type ProspectPhone struct {
	Personal string
	Work     string
}

type ProspectPersonal struct {
	Name       ProspectName
	Gender     string
	Occupation string
	Title      string
}

type ProspectName struct {
	First string
	Last  string
}

type ProspectSocial struct {
	Website  string
	Facebook string
	LinkedIn string
	Plus     string
	Quora    string
	Twitter  string
}

type ProspectMeta struct {
	OptOut bool `json: "opted_out"`
	Source string
	Notes  []string
	Tags   []string
	Custom []string
}

func (c *OutreachClient) GetProspect(id int) (Prospect, error) {
	p := Prospect{}

	if id == 0 {
		return p, errors.New("Prospect Id required.")
	}

	if c.Client == nil {
		return p, errors.New("You must assign an Outreach client.")
	}

	resp, e := c.Client.Get(fmt.Sprintf("https://api.outreach.io/1.0/prospects/%d", id))
	if e != nil {
		return p, e
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return p, e
	}

	e = json.Unmarshal(body, &p)
	if e != nil {
		return p, e
	}

	return p, nil
}

func (c *OutreachClient) PostProspect(p Prospect) error {
	if c.Client == nil {
		return errors.New("You must assign an Outreach client.")
	}

	return errors.New("Failed to post prospect")
}
