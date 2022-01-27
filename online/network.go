package online

import (
	"encoding/json"
	"time"
)

// DDoS Representation of an attack
type DDoS struct {
	ID               int       `json:"id"`
	Target           string    `json:"target"`
	StartDate        time.Time `json:"start"`
	EndDate          time.Time `json:"end"`
	MitigationSystem string    `json:"mitigation"`
	AttackType       string    `json:"type"`
}

// GetDDoS Retrieves the most recents DDoS alerts of an Online.net account
func GetDDoS() ([]DDoS, error) {
	var ddos []DDoS

	body, err := fetch("network/ddos")
	if err != nil {
		return ddos, err
	}

	err = json.Unmarshal(body, &ddos)
	if err != nil {
		return ddos, err
	}

	return ddos, nil
}
