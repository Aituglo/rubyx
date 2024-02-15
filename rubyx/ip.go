package rubyx

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aituglo/rubyx/pkg"
)

type IpResponse struct {
	Ips      []pkg.Ip `json:"ips"`
	TotalIps int      `json:"totalIps"`
}

func AddIp(config pkg.Config, ip string, program_id int) pkg.Ip {
	var data = []byte(fmt.Sprintf(`{
		"ip": "%s",
		"program_id": %d
	}`, ip, program_id))

	body := pkg.Post(config, "ip", data)

	var ipAdded pkg.Ip
	jsonErr := json.Unmarshal(body, &ipAdded)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return ipAdded
}

func GetAllIps(config pkg.Config) []pkg.Ip {
	body := pkg.Get(config, "ips?all=1")

	var ipResponse IpResponse
	jsonErr := json.Unmarshal(body, &ipResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return ipResponse.Ips
}

func GetIpsByProgram(config pkg.Config, program_id int) []pkg.Ip {
	body := pkg.Get(config, "ips?program_id="+fmt.Sprint(program_id))

	var ipResponse IpResponse
	jsonErr := json.Unmarshal(body, &ipResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return ipResponse.Ips
}
