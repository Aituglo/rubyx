package rubyx

import (
	"encoding/json"
	"log"

	"github.com/aituglo/rubyx/pkg"
)

func GetPlatforms(config pkg.Config) []pkg.Program {
	body := pkg.Get(config, "platform")

	var platforms []pkg.Program
	jsonErr := json.Unmarshal(body, &platforms)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return platforms
}
