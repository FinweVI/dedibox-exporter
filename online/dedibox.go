package online

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Product Represent an Offer (Dedibox)
type Product struct {
	ID     int         `json:"id"`
	Price  string      `json:"price"`
	Slug   string      `json:"slug"`
	Specs  ProductSpec `json:"specs"`
	Stocks []Stock     `json:"stocks"`
}

// ProductSpec Represent the specification of a Product
type ProductSpec struct {
	CPU          string `json:"cpu"`
	RAM          string `json:"ram"`
	Disks        string `json:"disks"`
	BP           string `json:"bp"`
	Customizable bool   `json:"customizable"`
	PolicerRate  string `json:"policer_rate"`
}

// Stock Represent the Stock of a Product
type Stock struct {
	Datacenter Datacenter `json:"datacenter"`
	Stock      int        `json:"stock"`
}

// Datacenter Represent a datacenter where a Product is located
type Datacenter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// GetPlans Retrieves the various offers and their availability
func GetPlans() ([]Product, error) {
	var plans []Product
	body, err := fetch("dedibox/plans")
	if err != nil {
		return plans, err
	}

	var ranges map[string]interface{}
	err = json.Unmarshal(body, &ranges)
	if err != nil {
		return plans, fmt.Errorf("Unable to unmarshal Dedibox Ranges into a Map")
	}

	for _, dediplans := range ranges {
		for _, details := range dediplans.(map[string]interface{}) {
			var p Product
			err = mapstructure.Decode(details, &p)
			if err != nil {
				return plans, fmt.Errorf("Unable to decode the Product data into the Product struct")
			}
			plans = append(plans, p)
		}
	}

	return plans, nil
}
