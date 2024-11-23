package api

import (
	"encoding/json"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type Torrent struct {
	ID       string   `json:"id"`
	Filename string   `json:"filename"`
	Hash     string   `json:"hash"`
	Bytes    int      `json:"bytes"`
	Host     string   `json:"host"`
	Split    int      `json:"split"`
	Progress float64  `json:"progress"`
	Status   string   `json:"status"`
	Added    string   `json:"added"`
	Links    []string `json:"links"`
	Ended    string   `json:"ended"`
	Speed    int      `json:"speed"`
	Seeders  int      `json:"seeders"`
}

type Torrents []*Torrent

func GetTorrentByHash(torrents Torrents, hash string) *Torrent {
	for _, torrent := range torrents {
		if torrent.Hash == hash {
			return torrent
		}
	}

	return nil
}

func GetTorrents(client *real_debrid.Client) (*Torrents, error) {
	url := client.GetURL("/torrents")

	query := url.Query()
	query.Add("limit", "1000")
	query.Add("page", "1")

	url.RawQuery = query.Encode()

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

	var torrents = &Torrents{}
	if err := json.NewDecoder(response.Body).Decode(torrents); err != nil {
		return nil, err
	}

	return torrents, nil
}
