package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	usage := `Rubyx
Usage:
  rubyx (new|rm|set|unset) <program>...
  rubyx program
  rubyx -h | --help
  rubyx --version
Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.ParseArgs(usage, nil, "Rubyx-CLI 1.0")

	readConfig()

	fmt.Println(arguments)
	fmt.Println(config)

	if arguments["new"] == true {
		fmt.Println("Creating a new program")
	}

}
