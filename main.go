package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aituglo/rubyx/pkg"
	"github.com/aituglo/rubyx/rubyx"
	"github.com/docopt/docopt-go"
)

var config pkg.Config
var inputData []string

func main() {
	usage := `Rubyx
	Usage:
		rubyx (new|rm|set|unset) <program>...
		rubyx programs
		rubyx subdomains [-p <program> | --all]
		rubyx subdomain (add) (-|<subdomain>...) [ -p <program> ]
		rubyx tool -t <tool>
		rubyx -h | --help
		rubyx --version
	Options:
		-h --help     Show this screen.
		-p <program>	Use the program
		--version     Show version.`

	arguments, _ := docopt.ParseArgs(usage, nil, "Rubyx-CLI 1.0")

	fmt.Println(arguments)

	pkg.ReadConfig(&config)

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
		rubyx.SetProgram(config, name)
	}

	if arguments["unset"] == true {
		config.ActiveProgram = ""
		pkg.WriteConfig(config)
	}

	if arguments["programs"] == true {
		programs := rubyx.GetAllPrograms(config)

		for _, program := range programs {
			fmt.Println(program.Slug)
		}
	}

	if arguments["new"] == true {
		name := arguments["<program>"].([]string)[0]
		rubyx.NewProgram(config, name)
	}

	if arguments["rm"] == true {
		name := arguments["<program>"].([]string)[0]
		rubyx.DeleteProgram(config, name)
	}

	if arguments["add"] == true && arguments["subdomain"] == true {
		var data []string
		var program_id int

		if arguments["-"] == true {
			data = inputData
		} else {
			data = arguments["<domain>"].([]string)
		}

		if arguments["-p"] != nil {
			program_id = rubyx.GetProgramID(config, arguments["-p"].(string))
		} else {
			program_id = rubyx.GetProgramID(config, config.ActiveProgram)
		}

		for _, domain := range data {

			var subdomain = []byte(fmt.Sprintf(`{
				"program_id": %d,
				"url": "%s"
			}`, program_id, domain))

			pkg.Post(config, "subdomain", subdomain)
		}

	}

	if arguments["subdomains"] == true {
		var program_id int
		var subdomains []pkg.Subdomain

		if arguments["--all"] == true {
			subdomains = rubyx.GetAllSubdomains(config)
		} else {
			if arguments["-p"] != nil {
				program_id = rubyx.GetProgramID(config, arguments["-p"].(string))
			} else {
				program_id = rubyx.GetProgramID(config, config.ActiveProgram)
			}
			subdomains = rubyx.GetSubdomainsByProgram(config, program_id)
		}

		for _, domain := range subdomains {
			fmt.Println(domain.Subdomain)
		}
	}

	if arguments["tool"] == true {
		var program_id int
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputData = append(inputData, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}

		if arguments["<tool>"].(string) == "wappago" {
			for _, line := range inputData {
				var parsed pkg.WappaGo
				err := json.Unmarshal([]byte(line), &parsed)
				if err != nil {
					log.Printf("Error unmarshalling JSON: %v\n", err)
				}

				if parsed.Infos.Screenshot != "" {

					filePath := "/tmp/screenshots/" + parsed.Infos.Screenshot

					fileContent, err := os.ReadFile(filePath)
					if err != nil {
						log.Printf("Error reading file: %v\n", err)
					}

					base64Encoded := base64.StdEncoding.EncodeToString(fileContent)
					parsed.Infos.Screenshot = base64Encoded

					err = os.Remove(filePath)
					if err != nil {
						log.Printf("Error deleting file: %v\n", err)
					}
				}

				domain, err := pkg.ExtractDomain(parsed.Url)
				if err != nil {
					log.Printf("Error when extracting domain: %v\n", err)
				}
				program_id = rubyx.GetProgramIDByScope(config, domain)
				var technologies string
				for _, technology := range parsed.Infos.Technologies {
					technologies += technology.Name + ","
				}

				if program_id != -1 {
					var subdomain = []byte(fmt.Sprintf(`{
						"program_id": %d,
						"url": "%s",
						"title": "%s",
						"technologies": "%s",
						"ip": "%s",
						"screenshot": "%s",
						"port": "%s",
						"content_length": %d,
						"status_code": %d
					}`, program_id, parsed.Url, parsed.Infos.Title, technologies, parsed.Infos.IP, parsed.Infos.Screenshot, strings.Join(parsed.Infos.Ports, ","), int32(parsed.Infos.ContentLength), int32(parsed.Infos.StatusCode)))
					pkg.Post(config, "subdomain", subdomain)
				}
			}
		}

	}

}
