package online

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// API Base URL
const API = "https://api.online.net/api/v1"

func fetch(urlPart string) ([]byte, error) {
	var output []byte

	u, err := url.Parse(fmt.Sprintf("%s/%s", API, urlPart))
	if err != nil {
		return output, fmt.Errorf("Unable to build URL: %s", urlPart)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return output, fmt.Errorf("Unable to build URL: %s", urlPart)
	}

	token := os.Getenv("ONLINE_API_TOKEN")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(req)
	if err != nil {
		return output, fmt.Errorf("Unable to perform request on '%s' part of Online.net's API", urlPart)
	}
	defer resp.Body.Close()

	output, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return output, fmt.Errorf("Unable to read the output returned when querying %s", urlPart)
	}

	return output, nil
}
