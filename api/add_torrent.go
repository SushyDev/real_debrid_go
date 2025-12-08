package api

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type addTorrentResponse struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

func AddTorrent(client *real_debrid.Client, torrent io.Reader) (*addTorrentResponse, error) {
	url := client.GetUrl("/torrents/addTorrent")

	// Build multipart form-data with field name "file"
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer writer.Close()

		part, err := writer.CreateFormFile("file", "upload.torrent")
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, torrent); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()

	req, err := http.NewRequest("PUT", url.String(), pr)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

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
