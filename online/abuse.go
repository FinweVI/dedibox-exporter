package online

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

// Abuse Representation from Online.net's API, limited to what we need
type Abuse struct {
	ID       uuid.UUID `json:"id"`
	Date     time.Time `json:"date"`
	Status   string    `json:"status"`
	Sender   string    `json:"sender"`
	Service  string    `json:"service"`
	Category string    `json:"type"`
}

// GetAbuses Retrieves all the pending abuses of an Online.net account
func GetAbuses() ([]Abuse, error) {
	var abuses []Abuse

	// We fetch only the first page of *unresolved* abuses
	// If you have more than one page... what are you doing on the Internet?
	body, err := fetch("abuse?status=pending")
	if err != nil {
		return abuses, err
	}

	err = json.Unmarshal(body, &abuses)
	if err != nil {
		return abuses, err
	}

	return abuses, nil
}
