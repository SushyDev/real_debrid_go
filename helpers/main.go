package helpers

import (
	"github.com/sushydev/real_debrid_go/api"
)

func GetTorrentByHash(torrents []*api.Torrent, hash string) *api.Torrent {
	for _, torrent := range torrents {
		if torrent.Hash == hash {
			return torrent
		}
	}

	return nil
}
