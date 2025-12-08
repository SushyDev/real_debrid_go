package api

import (
	"encoding/json"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type TrafficEntry struct {
	Left  int    `json:"left"`
	Bytes int    `json:"bytes"`
	Links int    `json:"links"`
	Limit int    `json:"limit"`
	Type  string `json:"type"`
	Extra int    `json:"extra"`
	Reset string `json:"reset"`
}

func GetTraffic(client *real_debrid.Client) (map[string]TrafficEntry, error) {
	url := client.GetUrl("/traffic")
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
	out := map[string]TrafficEntry{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

type TrafficDetailsDay struct {
	Host  map[string]int `json:"host"`
	Bytes int            `json:"bytes"`
}

func GetTrafficDetails(client *real_debrid.Client, start, end string) (map[string]TrafficDetailsDay, error) {
	url := client.GetUrl("/traffic/details")
	q := url.Query()
	if start != "" {
		q.Set("start", start)
	}
	if end != "" {
		q.Set("end", end)
	}
	url.RawQuery = q.Encode()
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
	out := map[string]TrafficDetailsDay{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}
