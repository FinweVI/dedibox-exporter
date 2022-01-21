package main

import (
	"fmt"
	"strings"
)

var validCollectors = []string{"abuse", "ddos", "plan", "dedibackup"}

type collectorSlice []string

func isValidCollector(c string) bool {
	for _, a := range validCollectors {
		if a == c {
			return true
		}
	}
	return false
}

func (cs *collectorSlice) Contains(val string) bool {
	for _, item := range *cs {
		if item == val {
			return true
		}
	}
	return false
}

func (cs *collectorSlice) String() string {
	return fmt.Sprintf("Collectors: %s", strings.Join(*cs, ", "))
}

func (cs *collectorSlice) Set(cltr string) error {
	if !cs.Contains(cltr) {
		if isValidCollector(cltr) {
			*cs = append(*cs, cltr)
		} else {
			return fmt.Errorf("Choose between: %s", strings.Join(validCollectors, ", "))
		}
	}

	return nil
}
