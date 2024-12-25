package api

import (
	"encoding/json"
	"net/http"
	urlpkg "net/url"
	"strings"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type UnrestrictLinkResponse struct {
	ID         string `json:"id"`
	Filename   string `json:"filename"`
	MimeType   string `json:"mimeType"`
	FileSize   int64  `json:"fileSize"`
	Link       string `json:"link"`
	Host       string `json:"host"`
	Chunks     int    `json:"chunks"`
	Crc        int    `json:"crc"`
	Download   string `json:"download"`
	Streamable int    `json:"streamable"`
}

func UnrestrictLink(client *real_debrid.Client, link string) (*UnrestrictLinkResponse, error) {
	url := client.GetURL("/unrestrict/link")

	form := urlpkg.Values{}
	form.Add("link", link)

	formReader := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", url.String(), formReader)
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

	var unrestrictLinkResponse UnrestrictLinkResponse
	err = json.NewDecoder(response.Body).Decode(&unrestrictLinkResponse)
	if err != nil {
		return nil, err
	}

	return &unrestrictLinkResponse, nil
}
