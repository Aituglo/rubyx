package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"regexp"
)

func ExtractDomain(input string) (string, error) {
	regex := regexp.MustCompile(`(?i)(?:https?://)?(?:\*\.)?([a-z0-9.-]+[a-z0-9])`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) < 2 {
		return "", fmt.Errorf("failed to extract domain from input: %s", input)
	}

	return matches[1], nil
}

func ReadConfig(config *Config) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	configFile, err := os.Open(usr.HomeDir + "/.rubyx/config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer configFile.Close()

	byteConfig, _ := io.ReadAll(configFile)

	json.Unmarshal(byteConfig, &config)
}

func WriteConfig(config Config) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	content, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(usr.HomeDir+"/.rubyx/config.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
