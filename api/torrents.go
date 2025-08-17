package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type Torrent struct {
	ID       string   `json:"id"`
	Filename string   `json:"filename"`
	Hash     string   `json:"hash"`
	Bytes    int      `json:"bytes"`
	Host     string   `json:"host"`
	Split    int      `json:"split"`
	Progress int      `json:"progress"`
	Status   string   `json:"status"`
	Added    string   `json:"added"`
	Links    []string `json:"links"`
	Ended    string   `json:"ended"`
	Speed    int      `json:"speed"`
	Seeders  int      `json:"seeders"`
}

type TorrentResponse struct {
	client        *real_debrid.Client
	torrents      []*Torrent
	currentPage   uint
	totalTorrents uint
}

func (response *TorrentResponse) GetTorrents() []*Torrent {
	return response.torrents
}

func (response *TorrentResponse) GetCurrentPage() uint {
	return response.currentPage
}

func (response *TorrentResponse) GetTotalTorrents() uint {
	return response.totalTorrents
}

func (response *TorrentResponse) NextPage(limit uint) (*TorrentResponse, error) {
	nextPage := response.currentPage + 1
	return GetTorrents(response.client, limit, nextPage)
}

func GetTorrents(client *real_debrid.Client, limit uint, page uint) (*TorrentResponse, error) {
	if limit < 1 || limit > 1000 {
		return nil, fmt.Errorf("limit must be between 1 and 1000, got %d", limit)
	}

	url := client.GetUrl("/torrents")

	query := url.Query()
	query.Add("limit", strconv.Itoa(int(limit)))
	query.Add("page", strconv.Itoa(int(page)))

	url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = client.HandleResponseCode(response, 200)
	if err != nil {
		return nil, err
	}

	var torrents = []*Torrent{}
	if err := json.NewDecoder(response.Body).Decode(&torrents); err != nil {
		return nil, err
	}

	totalCountHeader := response.Header.Get("X-Total-Count")
	if totalCountHeader == "" {
		return nil, fmt.Errorf("X-Total-Count header not found in response")
	}

	totalCount, err := strconv.Atoi(totalCountHeader)
	if err != nil {
		return nil, err
	}

	return &TorrentResponse{
		client:        client,
		currentPage:   page,
		torrents:      torrents,
		totalTorrents: uint(totalCount),
	}, nil
}

func GetAllTorrents(client *real_debrid.Client) ([]*Torrent, error) {
	// Fetch the first page to get the total count
	firstPage, err := GetTorrents(client, 1000, 1)
	if err != nil {
		return nil, err
	}

	totalPages := (firstPage.totalTorrents + 999) / 1000 // Calculate total pages based on 1000 limit
	allTorrents := make([]*Torrent, 0, firstPage.totalTorrents)

	// Fetch all pages
	for i := uint(1); i <= totalPages; i++ {
		page, err := GetTorrents(client, 1000, i)
		if err != nil {
			return nil, err
		}
		allTorrents = append(allTorrents, page.torrents...)
	}

	return allTorrents, nil
}
