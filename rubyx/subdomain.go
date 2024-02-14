package rubyx

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aituglo/rubyx/pkg"
)

type SubdomainResponse struct {
	Subdomains      []pkg.Subdomain `json:"subdomains"`
	TotalSubdomains int             `json:"totalSubdomains"`
}

func AddSubdomain(config pkg.Config, subdomain string) {
	program_id := GetProgramID(config, config.ActiveProgram)

	var data = []byte(fmt.Sprintf(`{
		"subdomain": "%s",
		"program_id": %d
	}`, subdomain, program_id))

	pkg.Post(config, "subdomain", data)
}

func GetAllSubdomains(config pkg.Config) []pkg.Subdomain {
	body := pkg.Get(config, "subdomains?all=1")

	var subdomainResponse SubdomainResponse
	jsonErr := json.Unmarshal(body, &subdomainResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return subdomainResponse.Subdomains
}

func GetSubdomainsByProgram(config pkg.Config, program_id int) []pkg.Subdomain {
	body := pkg.Get(config, "subdomains?program_id="+fmt.Sprint(program_id))

	var subdomainResponse SubdomainResponse
	jsonErr := json.Unmarshal(body, &subdomainResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return subdomainResponse.Subdomains
}
