package api

import (
	"encoding/json"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Points     int    `json:"points"`
	Locale     string `json:"locale,omitempty"`
	Avatar     string `json:"avatar"`
	Type       string `json:"type"`
	Premium    int    `json:"premium"`
	Expiration string `json:"expiration"`
}

func GetUser(client *real_debrid.Client) (*User, error) {
	url := client.GetUrl("/user")
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
	var u User
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
