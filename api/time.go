package api

import (
	"io"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

func GetTime(client *real_debrid.Client) (string, error) {
	url := client.GetUrl("/time")
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", client.HandleResponseCode(resp, 200)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GetTimeISO(client *real_debrid.Client) (string, error) {
	url := client.GetUrl("/time/iso")
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", client.HandleResponseCode(resp, 200)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
