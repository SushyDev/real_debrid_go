package api

import (
	"encoding/json"
	"net/http"
	urlpkg "net/url"
	"strings"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type Settings struct {
	DownloadPorts                []string          `json:"download_ports"`
	DownloadPort                 string            `json:"download_port"`
	Locales                      map[string]string `json:"locales"`
	Locale                       string            `json:"locale"`
	StreamingQualities           []string          `json:"streaming_qualities"`
	StreamingQuality             string            `json:"streaming_quality"`
	MobileStreamingQuality       string            `json:"mobile_streaming_quality"`
	StreamingLanguages           map[string]string `json:"streaming_languages"`
	StreamingLanguagePreference  string            `json:"streaming_language_preference"`
	StreamingCastAudio           []string          `json:"streaming_cast_audio"`
	StreamingCastAudioPreference string            `json:"streaming_cast_audio_preference"`
}

func GetSettings(client *real_debrid.Client) (*Settings, error) {
	url := client.GetUrl("/settings")
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
	var s Settings
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func UpdateSetting(client *real_debrid.Client, name, value string) error {
	url := client.GetUrl("/settings/update")
	form := urlpkg.Values{}
	form.Add("setting_name", name)
	form.Add("setting_value", value)
	req, err := http.NewRequest("POST", url.String(), strings.NewReader(form.Encode()))
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

func ConvertPoints(client *real_debrid.Client) error {
	url := client.GetUrl("/settings/convertPoints")
	req, err := http.NewRequest("POST", url.String(), nil)
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

func ChangePassword(client *real_debrid.Client) error {
	url := client.GetUrl("/settings/changePassword")
	req, err := http.NewRequest("POST", url.String(), nil)
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
