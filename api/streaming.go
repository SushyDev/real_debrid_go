package api

import (
	"encoding/json"
	"net/http"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type StreamingTranscode struct {
	Apple    map[string]string `json:"apple"`
	Dash     map[string]string `json:"dash"`
	LiveMP4  map[string]string `json:"liveMP4"`
	H264WebM map[string]string `json:"h264WebM"`
}

func GetStreamingTranscode(client *real_debrid.Client, id string) (*StreamingTranscode, error) {
	url := client.GetUrl("/streaming/transcode/" + id)
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
	var s StreamingTranscode
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

type MediaInfos struct {
	Filename string  `json:"filename"`
	Hoster   string  `json:"hoster"`
	Link     string  `json:"link"`
	Type     string  `json:"type"`
	Season   *string `json:"season"`
	Episode  *string `json:"episode"`
	Year     *string `json:"year"`
	Duration float64 `json:"duration"`
	Bitrate  int     `json:"bitrate"`
	Size     int     `json:"size"`
	Details  struct {
		Video map[string]struct {
			Stream     string `json:"stream"`
			Lang       string `json:"lang"`
			LangISO    string `json:"lang_iso"`
			Codec      string `json:"codec"`
			Colorspace string `json:"colorspace"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
		} `json:"video"`
		Audio map[string]struct {
			Stream   string  `json:"stream"`
			Lang     string  `json:"lang"`
			LangISO  string  `json:"lang_iso"`
			Codec    string  `json:"codec"`
			Sampling int     `json:"sampling"`
			Channels float64 `json:"channels"`
		} `json:"audio"`
		Subtitles map[string]struct {
			Stream  string `json:"stream"`
			Lang    string `json:"lang"`
			LangISO string `json:"lang_iso"`
			Type    string `json:"type"`
		} `json:"subtitles"`
	} `json:"details"`
	PosterPath   string `json:"poster_path"`
	AudioImage   string `json:"audio_image"`
	BackdropPath string `json:"backdrop_path"`
}

func GetStreamingMediaInfos(client *real_debrid.Client, id string) (*MediaInfos, error) {
	url := client.GetUrl("/streaming/mediaInfos/" + id)
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
	var m MediaInfos
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}
