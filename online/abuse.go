package online

import (
	"encoding/json"
	"fmt"
)

// Abuse Representation from Online.net's API
type Abuse struct {
	ID          int    `json:"id"`
	Date        string `json:"date"`
	Status      string `json:"status"`
	SentDate    string `json:"sent_date"`
	Description string `json:"description"`
	Sender      string `json:"sender"`
	Service     string `json:"service"`
	Category    string `json:"type"`
}

// GetAbuses Retrieves all the available abuses of an Online.net account
func GetAbuses() ([]Abuse, error) {
	var abuses []Abuse

	body, err := fetch("abuse")
	if err != nil {
		return abuses, err
	}

	err = json.Unmarshal(body, &abuses)
	if err != nil {
		return abuses, fmt.Errorf("Unable to unmarshal the returned JSON for abuses")
	}

	return abuses, nil
}
