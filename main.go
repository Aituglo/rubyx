package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/aituglo/rubyx-cli/pkg"
	"github.com/docopt/docopt-go"
)

var config pkg.Config

func readConfig() {
	usr, err := user.Current()

	configFile, err := os.Open(usr.HomeDir + "/.rubyx/config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer configFile.Close()

	byteConfig, _ := ioutil.ReadAll(configFile)

	json.Unmarshal(byteConfig, &config)
}

func writeConfig() {
	usr, err := user.Current()

	content, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(usr.HomeDir+"/.rubyx/config.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	usage := `Rubyx
Usage:
  rubyx (new|rm|set|unset) <program>...
  rubyx programs
  rubyx -h | --help
  rubyx --version
Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.ParseArgs(usage, nil, "Rubyx-CLI 1.0")

	readConfig()

	fmt.Println(arguments)

	if arguments["set"] == true {
		name := arguments["<program>"].([]string)[0]

		body := pkg.Get(config, "program/slug/"+name)

		var program pkg.Program
		jsonErr := json.Unmarshal(body, &program)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		config.ActiveProgram = program.Slug
		config.ActiveProgramID = int(program.ID)

		writeConfig()
	}

	if arguments["unset"] == true {
		config.ActiveProgram = ""
		config.ActiveProgramID = 0

		writeConfig()
	}

	if arguments["programs"] == true {
		body := pkg.Get(config, "program")

		var programs []pkg.Program
		jsonErr := json.Unmarshal(body, &programs)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for _, program := range programs {
			fmt.Println(program.Slug)
		}
	}

	if arguments["new"] == true {

		name := arguments["<program>"].([]string)[0]

		var data = []byte(fmt.Sprintf(`{
			"name": "%s",
			"slug": "%s",
			"type": "private"
		}`, name, name))

		pkg.Post(config, "program", data)

	}

	if arguments["rm"] == true {

		name := arguments["<program>"].([]string)[0]

		pkg.Delete(config, "program/slug/"+name)

	}

}
