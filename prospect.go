package outreach

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Prospect struct {
	Data   ProspectData `json:"data"`
	Errors []map[string]interface{}
}

type ProspectData struct {
	Attributes ProspectAttributes `json:"attributes"`
}

type ProspectAttributes struct {
	Address  ProspectAddress  `json:"address,omitempty"`
	Company  ProspectCompany  `json:"company,omitempty"`
	Contact  ProspectContact  `json:"contact"`
	Personal ProspectPersonal `json:"personal"`
	Social   ProspectSocial   `json:"social,omitempty"`
	Meta     ProspectMeta     `json:"metadata,omitempty"`
}

type ProspectAddress struct {
	City    string   `json:"city,omitempty"`
	State   string   `json:"state,omitempty"`
	Country string   `json:"country,omitempty"`
	Street  []string `json:"street,omitempty"`
	Zip     string   `json:"zip,omitempty"`
}

type ProspectCompany struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Industry string `json:"industry,omitempty"`
	Size     string `json:"size,omitempty"`
	Locality string `json:"locality,omitempty"`
}

type ProspectContact struct {
	Timezone string        `json:"timezone,omitempty"`
	Email    string        `json:"email"`
	Phone    ProspectPhone `json:"phone,omitempty"`
}

type ProspectPhone struct {
	Personal string `json:"personal,omitempty"`
	Work     string `json:"work,omitempty"`
}

type ProspectPersonal struct {
	Name       ProspectName `json:"name"`
	Gender     string       `json:"gender,omitempty"`
	Occupation string       `json:"occupation,omitempty"`
	Title      string       `json:"title,omitempty"`
}

type ProspectName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type ProspectSocial struct {
	Website  string `json:"website,omitempty"`
	Facebook string `json:"facebook,omitempty"`
	LinkedIn string `json:"linkedin,omitempty"`
	Plus     string `json:"plus,omitempty"`
	Quora    string `json:"quora,omitempty"`
	Twitter  string `json:"twitter,omitempty"`
}

type ProspectMeta struct {
	OptOut bool     `json:"opted_out"`
	Source string   `json:"source,omitempty"`
	Notes  []string `json:"notes,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	Custom []string `json:"custom,omitempty"`
}

type ProspectResponse struct {
	Data   ProspectResponseData
	Errors []map[string]interface{}
}

type ProspectResponseData struct {
	ID int
}

type ProspectInstance struct {
	Client *http.Client
}

func (p *Prospect) Read(t []byte) (int, error) {
	j, e := json.Marshal(p)
	if e != nil {
		return 0, e
	}

	if len(t) < len(j) {
		return len(t), nil
	}

	copy(t, j)
	return len(j), io.EOF
}

func (p *Prospect) Close() error {
	return nil
}

func (i *ProspectInstance) Get(id int) (Prospect, error) {
	p := Prospect{}

	if i.Client == nil {
		return p, errors.New("You must assign a HTTP client.")
	}

	resp, e := i.Client.Get("https://api.outreach.io/1.0/prospect/" + strconv.Itoa(id))
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

	if len(p.Errors) > 0 {
		return p, fmt.Errorf("Got error response: %+v\n", p.Errors)
	}

	return p, nil
}

func (i *ProspectInstance) Post(p Prospect) (ProspectResponse, error) {
	pr := ProspectResponse{}

	if i.Client == nil {
		return pr, errors.New("You must assign a HTTP client.")
	}

	resp, e := i.Client.Post("https://api.outreach.io/1.0/prospect", "application/json", &p)
	if e != nil {
		return pr, e
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return pr, e
	}

	e = json.Unmarshal(body, &pr)
	if e != nil {
		return pr, e
	}

	if len(pr.Errors) > 0 {
		return pr, fmt.Errorf("Got error response: %+v\n", pr.Errors)
	}

	return pr, nil
}
