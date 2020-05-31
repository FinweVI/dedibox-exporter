package online

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// API Base URL
const API = "https://api.online.net/api/v1"

func fetch(urlPart string) ([]byte, error) {
	var output []byte

	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", API, urlPart)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return output, fmt.Errorf("Unable to build URL for %s", urlPart)
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
