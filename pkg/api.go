package pkg

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

	body, readErr := ioutil.ReadAll(res.Body)
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

	body, readErr := ioutil.ReadAll(res.Body)
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

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
