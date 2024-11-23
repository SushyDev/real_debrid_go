package api

import (
	"bytes"
	"encoding/json"
	real_debrid "github.com/sushydev/real_debrid_go"
	"net/http"
	"net/url"
)

type addMagnetResponse struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

func AddMagnet(client *real_debrid.Client, magnet string) (*addMagnetResponse, error) {
	input := url.Values{}
	input.Set("magnet", magnet)
	requestBody := input.Encode()
	reader := bytes.NewBufferString(requestBody)

	url := client.GetURL("/torrents/addMagnet")

	req, err := http.NewRequest("POST", url.String(), reader)
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

	data := &addMagnetResponse{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
