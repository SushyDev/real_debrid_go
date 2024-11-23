package api

import (
	"bytes"
	"net/http"
	urlpkg "net/url"

	real_debrid "github.com/sushydev/real_debrid_go"
)

func SelectFiles(client *real_debrid.Client, torrentId string, fileIds string) error {
	var input = urlpkg.Values{}
	input.Set("files", fileIds)

	requestBody := input.Encode()
	reader := bytes.NewBufferString(requestBody)

	url := client.GetURL("/torrents/selectFiles/" + torrentId)

	req, err := http.NewRequest("POST", url.String(), reader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = client.HandleResponseCode(response, 204)
	if err != nil {
		return err
	}

	return nil
}
