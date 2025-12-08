package api

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	urlpkg "net/url"
	"strings"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type ActiveCount struct {
	Nb    int `json:"nb"`
	Limit int `json:"limit"`
}

type AvailableHost struct {
	Host        string `json:"host"`
	MaxFileSize int    `json:"max_file_size"`
}

func GetTorrentByHash(torrents []*Torrent, hash string) *Torrent {
	for _, torrent := range torrents {
		if torrent.Hash == hash {
			return torrent
		}
	}

	return nil
}

func GetActiveCount(client *real_debrid.Client) (*ActiveCount, error) {
	url := client.GetUrl("/torrents/activeCount")
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 200); err != nil {
		return nil, err
	}
	var ac ActiveCount
	if err := json.NewDecoder(resp.Body).Decode(&ac); err != nil {
		return nil, err
	}
	return &ac, nil
}

func GetAvailableHosts(client *real_debrid.Client) ([]AvailableHost, error) {
	url := client.GetUrl("/torrents/availableHosts")
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 200); err != nil {
		return nil, err
	}
	var hosts []AvailableHost
	if err := json.NewDecoder(resp.Body).Decode(&hosts); err != nil {
		return nil, err
	}
	return hosts, nil
}

// AddTorrentMultipart uploads a torrent file using multipart and optional host
func AddTorrentMultipart(client *real_debrid.Client, data []byte, filename string, host string) (*addTorrentResponse, error) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	fw, err := w.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err := fw.Write(data); err != nil {
		return nil, err
	}
	if host != "" {
		w.WriteField("host", host)
	}
	w.Close()
	url := client.GetUrl("/torrents/addTorrent")
	req, err := http.NewRequest("PUT", url.String(), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 201); err != nil {
		return nil, err
	}
	var out addTorrentResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AddMagnetWithHost supports optional host
func AddMagnetWithHost(client *real_debrid.Client, magnet string, host string) (*addMagnetResponse, error) {
	form := urlpkg.Values{}
	form.Add("magnet", magnet)
	if host != "" {
		form.Add("host", host)
	}
	url := client.GetUrl("/torrents/addMagnet")
	req, err := http.NewRequest("POST", url.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 201); err != nil {
		return nil, err
	}
	var out addMagnetResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
