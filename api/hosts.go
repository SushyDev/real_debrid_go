package api

import (
	"encoding/json"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type HostBasic struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func GetHosts(client *real_debrid.Client) (map[string]HostBasic, error) {
	url := client.GetUrl("/hosts")
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
	out := map[string]HostBasic{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

type HostStatus struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	Supported         int    `json:"supported"`
	Status            string `json:"status"`
	CheckTime         string `json:"check_time"`
	CompetitorsStatus map[string]struct {
		Status    string `json:"status"`
		CheckTime string `json:"check_time"`
	} `json:"competitors_status"`
}

func GetHostsStatus(client *real_debrid.Client) (map[string]HostStatus, error) {
	url := client.GetUrl("/hosts/status")
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
	out := map[string]HostStatus{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

func GetHostsRegex(client *real_debrid.Client) ([]string, error) {
	url := client.GetUrl("/hosts/regex")
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
	var out []string
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

func GetHostsRegexFolder(client *real_debrid.Client) ([]string, error) {
	url := client.GetUrl("/hosts/regexFolder")
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
	var out []string
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

func GetHostsDomains(client *real_debrid.Client) ([]string, error) {
	url := client.GetUrl("/hosts/domains")
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
	var out []string
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}
