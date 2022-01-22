package online

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Server Representation of a server (it's ID)
type Server struct {
	ID int `json:"id"`
}

// Dedibackup Representation of a Dedibackup
type Dedibackup struct {
	Login          string `json:"login"`
	Server         string `json:"server"`
	Active         bool   `json:"active"`
	ACLEnabled     bool   `json:"acl_enabled"`
	Autologin      bool   `json:"autologin"`
	QuotaSpace     int64  `json:"quota_space"`
	QuotaSpaceUsed int64  `json:"quota_space_used"`
	QuotaFiles     int    `json:"quota_files"`
	QuotaFilesUsed int    `json:"quota_files_used"`
}

func getServers() ([]Server, error) {
	var servers []Server

	body, err := fetch("server")
	if err != nil {
		return servers, err
	}

	var tmpData []string
	err = json.Unmarshal(body, &tmpData)
	if err != nil {
		return servers, err
	}

	for _, serverLink := range tmpData {
		splt := strings.Split(serverLink, "/")
		sid, err := strconv.Atoi(splt[len(splt)-1])
		if err != nil {
			return servers, err
		}
		servers = append(servers, Server{ID: sid})
	}

	return servers, nil
}

func getDedibackup(serverID int) (Dedibackup, error) {
	var dedibackup Dedibackup

	body, err := fetch(fmt.Sprintf("server/backup/%s", strconv.Itoa(serverID)))
	if err != nil {
		return dedibackup, err
	}

	err = json.Unmarshal(body, &dedibackup)
	if err != nil {
		return dedibackup, err
	}

	return dedibackup, nil
}

// GetDedibackups Retrieves all the dedibackups and their statuses
func GetDedibackups() ([]Dedibackup, error) {
	var dedibackups []Dedibackup

	srvs, err := getServers()
	if err != nil {
		return dedibackups, err
	}

	var wg sync.WaitGroup
	results := make(chan Dedibackup, 2)
	errors := make(chan error, 2)
	for _, srv := range srvs {
		// call getDedibackup as a gorouting and wait later the result
		wg.Add(1)
		go func(sid int) {
			defer wg.Done()
			dedibkp, err := getDedibackup(sid)
			if err != nil {
				errors <- err
				return
			}
			results <- dedibkp
		}(srv.ID)
	}

	wg.Wait()
	close(results)
	close(errors)

	for err := range errors {
		return dedibackups, err
	}
	for res := range results {
		dedibackups = append(dedibackups, res)
	}

	return dedibackups, nil
}
