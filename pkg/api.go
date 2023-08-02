package pkg

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetProgramID(config Config, name string) int {
	body := Get(config, "program/slug/"+name)

	var program Program
	jsonErr := json.Unmarshal(body, &program)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return int(program.ID)
}

func GetProgramIDByScope(config Config, scope string) int {
	body := Get(config, "scope?subdomain="+scope)

	program_id, err := strconv.Atoi(strings.TrimSpace(string(body)))
	if err != nil {
		fmt.Println("Error when converting: ", err)
	}

	return program_id
}

func Get(config Config, endpoint string) []byte {
	url := config.Url + endpoint

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cookie", "rubyx-jwt="+config.ApiKey)

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode == 401 {
		log.Fatal("Error in authentication, please verify api key")
	}

	if res.StatusCode != 200 {
		log.Fatal("An error has occured")
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}

func Post(config Config, endpoint string, data []byte) []byte {
	url := config.Url + endpoint

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cookie", "rubyx-jwt="+config.ApiKey)

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode == 401 {
		log.Fatal("Error in authentication, please verify api key")
	}

	if res.StatusCode != 200 {
		log.Fatal("An error has occured")
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(string(body))

	return body
}

func Delete(config Config, endpoint string) []byte {
	url := config.Url + endpoint

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cookie", "rubyx-jwt="+config.ApiKey)

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode == 401 {
		log.Fatal("Error in authentication, please verify api key")
	}

	if res.StatusCode != 200 {
		log.Fatal("An error has occured")
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
