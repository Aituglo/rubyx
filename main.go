package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aituglo/rubyx/pkg"
	"github.com/aituglo/rubyx/rubyx"
	"github.com/aituglo/rubyx/tool"
	"github.com/docopt/docopt-go"
)

var config pkg.Config
var inputData []string

func main() {
	usage := `Rubyx
	Usage:
		rubyx (new|rm|set|unset) <program>...
		rubyx current
		rubyx programs [reload]
		rubyx scope [-p <program>]
		rubyx subdomains [-p <program> | --all]
		rubyx subdomain (add) (-|<subdomain>...) [ -p <program> ]
		rubyx ips [-p <program> | --all]
		rubyx technologies
		rubyx tool -t <tool>
		rubyx -h | --help
		rubyx --version
	Options:
		-h --help     Show this screen.
		-p <program>	Use the program
		--version     Show version.`

	arguments, _ := docopt.ParseArgs(usage, nil, "Rubyx-CLI 1.0")

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

	if arguments["current"] == true {
		fmt.Println(config.ActiveProgram)
	}

	if arguments["unset"] == true {
		config.ActiveProgram = ""
		pkg.WriteConfig(config)
	}

	if arguments["programs"] == true {
		if arguments["reload"] == true {
			rubyx.ReloadPrograms(config)
		}
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
			data = arguments["<subdomain>"].([]string)
		}

		if arguments["-p"] != nil {
			program_id = rubyx.GetProgramID(config, arguments["-p"].(string))
		} else {
			program_id = rubyx.GetProgramID(config, config.ActiveProgram)
		}

		for _, domain := range data {
			rubyx.AddSubdomain(config, domain, program_id)
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

	if arguments["ips"] == true {
		var program_id int
		var ips []pkg.Ip

		if arguments["--all"] == true {
			ips = rubyx.GetAllIps(config)
		} else {
			if arguments["-p"] != nil {
				program_id = rubyx.GetProgramID(config, arguments["-p"].(string))
			} else {
				program_id = rubyx.GetProgramID(config, config.ActiveProgram)
			}
			ips = rubyx.GetIpsByProgram(config, program_id)
		}

		for _, ip := range ips {
			fmt.Println(ip.Ip)
		}
	}

	if arguments["scope"] == true {
		var program_id int
		var scope []pkg.Scope

		if arguments["-p"] != nil {
			program_id = rubyx.GetProgramID(config, arguments["-p"].(string))
		} else {
			program_id = rubyx.GetProgramID(config, config.ActiveProgram)
		}
		scope = rubyx.GetScope(config, program_id)

		for _, domain := range scope {
			fmt.Println(domain.Scope)
		}
	}

	if arguments["technologies"] == true {
		technologies := rubyx.GetAllTechnologies(config)

		for _, technology := range technologies {
			fmt.Println(technology.Name)
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

		if arguments["-p"] != nil {
			program_id = rubyx.GetProgramID(config, arguments["-p"].(string))
		} else {
			program_id = rubyx.GetProgramID(config, config.ActiveProgram)
		}

		if arguments["<tool>"].(string) == "wappago" {
			tool.WappaGo(config, inputData, program_id)
		}

	}

}
