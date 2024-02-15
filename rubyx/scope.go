package rubyx

import (
	"encoding/json"
	"fmt"

	"github.com/aituglo/rubyx/pkg"
)

func GetScope(config pkg.Config, program_id int) []pkg.Scope {
	body := pkg.Get(config, "scope/"+fmt.Sprint(program_id))

	var scopes []pkg.Scope
	jsonErr := json.Unmarshal(body, &scopes)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return scopes
}
