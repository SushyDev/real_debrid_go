package real_debrid_go

import (
	"fmt"
	"net/http"
	urlpkg "net/url"
)

type Client struct {
	http.Client
	token string

	host string
	path string
}

func NewClient(token string, client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}

	return &Client{
		Client: *client,
		token:  token,

		host: "https://api.real-debrid.com",
		path: "/rest/1.0",
	}
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.token))

	return client.Client.Do(req)
}

// --- Helpers

func (client *Client) GetUrl(endpoint string) *urlpkg.URL {
	url, _ := urlpkg.Parse(client.host)
	url.Path += client.path + endpoint

	return url
}

func (client *Client) HandleResponseCode(response *http.Response, expected int) error {
	switch response.StatusCode {
	case expected:
		return nil
	case 204:
		return fmt.Errorf("No content")
	case 400:
		return fmt.Errorf("Bad Request (see error message)")
	case 401:
		return fmt.Errorf("Bad token (expired, invalid)")
	case 403:
		return fmt.Errorf("Permission denied (account locked, not premium) or Infringing torrent")
	case 503:
		return fmt.Errorf("Service unavailable (see error message)")
	case 504:
		return fmt.Errorf("Service timeout (see error message)")
	default:
		return fmt.Errorf("[%v] Unknown error", response.StatusCode)
	}
}
