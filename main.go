package main

import (
	"bufio"
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
var inputData []string

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
	rubyx domains [-p <program> | --all]
	rubyx domain (add) (-|<domain>...) [ -p <program> ]
  rubyx -h | --help
  rubyx --version
Options:
  -h --help     Show this screen.
	-p <program>	Use the program
  --version     Show version.`

	arguments, _ := docopt.ParseArgs(usage, nil, "Rubyx-CLI 1.0")

	readConfig()

	fmt.Println(arguments)

	if arguments["-"] == true {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputData = append(inputData, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}

	if arguments["set"] == true {
		name := arguments["<program>"].([]string)[0]

		body := pkg.Get(config, "program/slug/"+name)

		var program pkg.Program
		jsonErr := json.Unmarshal(body, &program)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		config.ActiveProgram = program.Slug

		writeConfig()
	}

	if arguments["unset"] == true {
		config.ActiveProgram = ""

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

	if arguments["add"] == true && arguments["domain"] == true {
		var data []string
		var program_id int

		if arguments["-"] == true {
			data = inputData
		} else {
			data = arguments["<domain>"].([]string)
		}

		if arguments["-p"] != nil {
			program_id = pkg.GetProgramID(config, arguments["-p"].(string))
		} else {
			program_id = pkg.GetProgramID(config, config.ActiveProgram)
		}

		for _, domain := range data {

			var subdomain = []byte(fmt.Sprintf(`{
				"program_id": %d,
				"url": "%s"
			}`, program_id, domain))

			pkg.Post(config, "subdomain", subdomain)
		}

	}

	if arguments["domains"] == true {
		var body []byte

		if arguments["-p"] != nil {
			program_id := pkg.GetProgramID(config, arguments["-p"].(string))

			body = pkg.Get(config, "subdomain/program/"+fmt.Sprint(program_id))

		} else if arguments["--all"] == true {
			body = pkg.Get(config, "subdomain")
		} else {
			program_id := pkg.GetProgramID(config, config.ActiveProgram)

			body = pkg.Get(config, "subdomain/program/"+fmt.Sprint(program_id))
		}

		var domains []pkg.Subdomain
		jsonErr := json.Unmarshal(body, &domains)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for _, domain := range domains {
			fmt.Println(domain.Url)
		}
	}

}
