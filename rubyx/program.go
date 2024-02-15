package rubyx

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aituglo/rubyx/pkg"
)

func SetProgram(config pkg.Config, name string) {
	body := pkg.Get(config, "program/"+name)

	var program pkg.Program
	jsonErr := json.Unmarshal(body, &program)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	config.ActiveProgram = program.Slug
	pkg.WriteConfig(config)
}

func DeleteProgram(config pkg.Config, name string) {
	id := GetProgramID(config, name)

	pkg.Delete(config, "program/"+strconv.Itoa(id))
}

func NewProgram(config pkg.Config, name string) {
	platform_id := 0
	platforms := GetPlatforms(config)
	for _, platform := range platforms {
		if platform.Slug == "default" {
			platform_id = platform.ID
		}
	}

	var data = []byte(fmt.Sprintf(`{
		"platform_id": %d,
		"name": "%s",
		"slug": "%s",
		"type": "private"
	}`, platform_id, name, name))

	pkg.Post(config, "program", data)
}

func ReloadPrograms(config pkg.Config) {
	pkg.Get(config, "programs?reload=1")
}

func GetAllPrograms(config pkg.Config) []pkg.Program {
	body := pkg.Get(config, "programs?all=1")

	var programs []pkg.Program
	jsonErr := json.Unmarshal(body, &programs)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return programs
}

func GetProgramID(config pkg.Config, name string) int {
	body := pkg.Get(config, "program/"+name)

	var program pkg.Program
	jsonErr := json.Unmarshal(body, &program)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return int(program.ID)
}

func GetProgramIDByScope(config pkg.Config, scope string) int {
	body := pkg.Get(config, "scope?subdomain="+scope)

	program_id, err := strconv.Atoi(strings.TrimSpace(string(body)))
	if err != nil {
		fmt.Println("Error when converting: ", err)
	}

	return program_id
}
