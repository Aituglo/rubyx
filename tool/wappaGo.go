package tool

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aituglo/rubyx/pkg"
	"github.com/aituglo/rubyx/rubyx"
)

func WappaGo(config pkg.Config, inputData []string, program_id int) {
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

		ip_id := rubyx.AddIp(config, parsed.Infos.IP, program_id).ID

		for _, port := range parsed.Infos.Ports {
			portInt, err := strconv.Atoi(port)
			if err != nil {
				log.Printf("Error converting port to integer: %v\n", err)
				continue
			}
			rubyx.AddPort(config, portInt, int(ip_id))
		}

		fmt.Printf(`{
			"program_id": %d,
			"subdomain": "%s",
			"title": "%s",
			"ip": %d,
			"screenshot": "%s",
			"content_length": %d,
			"status_code": %d,
			"tag": "wappago"
		}`, program_id, domain, parsed.Infos.Title, ip_id, parsed.Infos.Screenshot, int32(parsed.Infos.ContentLength), int32(parsed.Infos.StatusCode))

		var subdomain = []byte(fmt.Sprintf(`{
				"program_id": %d,
				"subdomain": "%s",
				"title": "%s",
				"ip": %d,
				"screenshot": "%s",
				"content_length": %d,
				"status_code": %d,
				"tag": "wappago"
			}`, program_id, domain, parsed.Infos.Title, ip_id, parsed.Infos.Screenshot, int32(parsed.Infos.ContentLength), int32(parsed.Infos.StatusCode)))

		subdomainResponse := pkg.Post(config, "subdomain", subdomain)
		var addedSubdomain pkg.Subdomain
		err = json.Unmarshal([]byte(subdomainResponse), &addedSubdomain)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v\n", err)
		}

		for _, technology := range parsed.Infos.Technologies {
			rubyx.AddTechnologie(config, addedSubdomain.ID, technology.Name, technology.Version)
		}
	}
}
