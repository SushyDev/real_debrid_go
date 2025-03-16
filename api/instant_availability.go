package api

import (
	"encoding/json"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type InstantAvailabilityFile struct {
	FileName string `json:"filename"`
	FileSize int    `json:"filesize"`
}

type InstantAvailability struct {
	hash  string
	files map[string]map[string][]map[string]InstantAvailabilityFile
}

func GetInstantAvailability(client *real_debrid.Client, hash string) (*InstantAvailability, error) {
	url := client.GetUrl("/torrents/instantAvailability/" + hash)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = client.HandleResponseCode(response, 200)
	if err != nil {
		return nil, err
	}

	files := make(map[string]map[string][]map[string]InstantAvailabilityFile)

	err = json.NewDecoder(response.Body).Decode(&files)
	if err != nil {
		return nil, err
	}

	instantAvailability := &InstantAvailability{
		hash:  hash,
		files: files,
	}

	return instantAvailability, nil
}

func (instantAvailability *InstantAvailability) GetFiles() map[string]map[string][]map[string]InstantAvailabilityFile {
	return instantAvailability.files
}

func (instantAvailability *InstantAvailability) GetHash() string {
	return instantAvailability.hash
}
