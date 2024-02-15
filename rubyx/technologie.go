package rubyx

import (
	"encoding/json"
	"fmt"

	"github.com/aituglo/rubyx/pkg"
)

func AddTechnologie(config pkg.Config, subdomain_id int64, name string, version string) {
	var data = []byte(fmt.Sprintf(`{
		"subdomain_id": %d,
		"name": "%s",
		"version": "%s"
	}`, subdomain_id, name, version))

	pkg.Post(config, "technologie", data)
}

func GetAllTechnologies(config pkg.Config) []pkg.Technology {
	body := pkg.Get(config, "technologies")

	var technologies []pkg.Technology
	jsonErr := json.Unmarshal(body, &technologies)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return technologies
}
