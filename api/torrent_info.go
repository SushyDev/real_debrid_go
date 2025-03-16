package api

import (
	"encoding/json"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type TorrentFile struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Bytes    int    `json:"bytes"`
	Selected int    `json:"selected"`
}

type TorrentInfo struct {
	ID               string        `json:"id"`
	Filename         string        `json:"filename"`
	OriginalFilename string        `json:"original_filename"`
	Hash             string        `json:"hash"`
	Bytes            int           `json:"bytes"`
	OriginalBytes    int           `json:"original_bytes"`
	Host             string        `json:"host"`
	Split            int           `json:"split"`
	Progress         float64       `json:"progress"`
	Status           string        `json:"status"`
	Added            string        `json:"added"`
	Files            []TorrentFile `json:"files"`
	Links            []string      `json:"links"`
	Ended            string        `json:"ended,omitempty"`
	Speed            int           `json:"speed,omitempty"`
	Seeders          int           `json:"seeders,omitempty"`
}

func GetTorrentInfo(client *real_debrid.Client, id string) (*TorrentInfo, error) {
	url := client.GetUrl("/torrents/info/" + id)

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

	torrentInfo := &TorrentInfo{}

	err = json.NewDecoder(response.Body).Decode(torrentInfo)
	if err != nil {
		return nil, err
	}

	return torrentInfo, nil
}
