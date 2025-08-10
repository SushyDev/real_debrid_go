package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type Download struct {
	ID        string `json:"id"`
	Filename  string `json:"filename"`
	MimeType  string `json:"mimeType"`
	Filesize  int64  `json:"filesize"`
	Link      string `json:"link"`
	Host      string `json:"host"`
	Chunks    int    `json:"chunks"`
	Crc       int    `json:"crc"`
	Download  string `json:"download"`
	Generated string `json:"generated"`
	Type      string `json:"type,omitempty"`
}

// GetDownloads returns downloads list with total count from X-Total-Count header
func GetDownloads(client *real_debrid.Client, limit uint, page uint, offset *uint) ([]*Download, int, error) {
	url := client.GetUrl("/downloads")
	q := url.Query()
	if limit > 0 {
		q.Add("limit", strconv.Itoa(int(limit)))
	}
	if page > 0 {
		q.Add("page", strconv.Itoa(int(page)))
	}
	if offset != nil {
		q.Add("offset", strconv.Itoa(int(*offset)))
	}
	url.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 200); err != nil {
		return nil, 0, err
	}
	var items []*Download
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, 0, err
	}
	totalCountHeader := resp.Header.Get("X-Total-Count")
	if totalCountHeader == "" {
		return items, 0, nil
	}
	totalCount, err := strconv.Atoi(totalCountHeader)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid X-Total-Count: %w", err)
	}
	return items, totalCount, nil
}

func DeleteDownload(client *real_debrid.Client, id string) error {
	url := client.GetUrl("/downloads/delete/" + id)
	req, err := http.NewRequest("DELETE", url.String(), nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return client.HandleResponseCode(resp, 204)
}
