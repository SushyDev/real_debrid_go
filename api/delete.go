package api

import (
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

func Delete(client *real_debrid.Client, id string) error {
	url := client.GetUrl("/torrents/delete/" + id)

	req, err := http.NewRequest("DELETE", url.String(), nil)
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
