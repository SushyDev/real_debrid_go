package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"

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

func GetTorrentByHash(torrents []*Torrent, hash string) *Torrent {
	for _, torrent := range torrents {
		if torrent.Hash == hash {
			return torrent
		}
	}

	return nil
}

func GetTorrents(client *real_debrid.Client, limit uint, page uint) ([]*Torrent, int, error) {
	url := client.GetUrl("/torrents")

	query := url.Query()
	query.Add("limit", strconv.Itoa(int(limit)))
	query.Add("page", strconv.Itoa(int(page)))

	url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, 0, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer response.Body.Close()

	err = client.HandleResponseCode(response, 200)
	if err != nil {
		return nil, 0, err
	}

	var torrents = []*Torrent{}
	if err := json.NewDecoder(response.Body).Decode(&torrents); err != nil {
		return nil, 0, err
	}

	totalCountHeader := response.Header.Get("X-Total-Count")
	if totalCountHeader == "" {
		return nil, 0, fmt.Errorf("X-Total-Count header not found in response")
	}

	totalCount, err := strconv.Atoi(totalCountHeader)
	if err != nil {
		return nil, 0, err
	}

	return torrents, totalCount, nil
}

