package api

import (
	"encoding/json"
	"io"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type addTorrentResponse struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

func AddTorrent(client *real_debrid.Client, torrent io.Reader) (*addTorrentResponse, error) {
	url := client.GetURL("/torrents/addTorrent")

	req, err := http.NewRequest("PUT", url.String(), torrent)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = client.HandleResponseCode(response, 201)
	if err != nil {
		return nil, err
	}

	data := &addTorrentResponse{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
