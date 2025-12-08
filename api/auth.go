package api

import (
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

// DisableAccessToken disables current access token and expects 204
func DisableAccessToken(client *real_debrid.Client) error {
	url := client.GetUrl("/disable_access_token")
	req, err := http.NewRequest("GET", url.String(), nil)
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
