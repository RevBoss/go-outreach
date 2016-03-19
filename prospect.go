package outreach

import (
	"errors"
	"net/http"
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

type ProspectInstance struct {
	Client *http.Client
}

func (i *ProspectInstance) Get(id int) (Prospect, error) {
	p := Prospect{}

	return p, nil
}

func (i *ProspectInstance) Post(p Prospect) error {
	if i.Client == nil {
		return errors.New("You must assign a HTTP client.")
	}

	return errors.New("Failed to post prospect")
}
