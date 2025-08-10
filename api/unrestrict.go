package api

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	urlpkg "net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	real_debrid "github.com/sushydev/real_debrid_go"
)

type UnrestrictCheckResponse struct {
	Host      string `json:"host"`
	Link      string `json:"link"`
	Filename  string `json:"filename"`
	Filesize  int64  `json:"filesize"`
	Supported int    `json:"supported"`
}

type UnrestrictFolderItem struct {
	ID        string `json:"id"`
	Filename  string `json:"filename"`
	MimeType  string `json:"mimeType"`
	Filesize  int64  `json:"filesize"`
	Link      string `json:"link"`
	Host      string `json:"host"`
	Chunks    int    `json:"chunks"`
	Download  string `json:"download"`
	Generated string `json:"generated"`
	Type      string `json:"type,omitempty"`
}

// UnrestrictLinkOptions are optional parameters to /unrestrict/link
type UnrestrictLinkOptions struct {
	Password string
	Remote   *int // 0 or 1
}

// UnrestrictLinkEnhanced supports optional password/remote and multiple response shapes
func UnrestrictLinkEnhanced(client *real_debrid.Client, link string, opts *UnrestrictLinkOptions) (*UnrestrictLinkResponse, []UnrestrictFolderItem, error) {
	url := client.GetUrl("/unrestrict/link")
	form := urlpkg.Values{}
	form.Add("link", link)
	if opts != nil {
		if opts.Password != "" {
			form.Add("password", opts.Password)
		}
		if opts.Remote != nil {
			form.Add("remote", strconv.Itoa(*opts.Remote))
		}
	}
	req, err := http.NewRequest("POST", url.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 200); err != nil {
		return nil, nil, err
	}
	dec := json.NewDecoder(resp.Body)
	// Try single object first
	var single UnrestrictLinkResponse
	if err := dec.Decode(&single); err == nil && single.Download != "" {
		return &single, nil, nil
	}
	// If not single, decode as array
	resp.Body.Close()
	// Re-do request to read body again (can't rewind easily)
	req, _ = http.NewRequest("POST", url.String(), strings.NewReader(form.Encode()))
	resp, err = client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 200); err != nil {
		return nil, nil, err
	}
	var multi []UnrestrictFolderItem
	if err := json.NewDecoder(resp.Body).Decode(&multi); err != nil {
		return nil, nil, err
	}
	return nil, multi, nil
}

func UnrestrictCheck(client *real_debrid.Client, link string, password string) (*UnrestrictCheckResponse, error) {
	url := client.GetUrl("/unrestrict/check")
	form := urlpkg.Values{}
	form.Add("link", link)
	if password != "" {
		form.Add("password", password)
	}
	req, err := http.NewRequest("POST", url.String(), strings.NewReader(form.Encode()))
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
	var r UnrestrictCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func UnrestrictFolder(client *real_debrid.Client, link string) ([]UnrestrictFolderItem, error) {
	url := client.GetUrl("/unrestrict/folder")
	form := urlpkg.Values{}
	form.Add("link", link)
	req, err := http.NewRequest("POST", url.String(), strings.NewReader(form.Encode()))
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
	var items []UnrestrictFolderItem
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, err
	}
	return items, nil
}

// UnrestrictContainerFile uploads a container file (RSDF/CCF/DLC)
func UnrestrictContainerFile(client *real_debrid.Client, filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	part, err := w.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, f); err != nil {
		return nil, err
	}
	w.Close()
	url := client.GetUrl("/unrestrict/containerFile")
	req, err := http.NewRequest("PUT", url.String(), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := client.HandleResponseCode(resp, 200); err != nil {
		return nil, err
	}
	var links []string
	if err := json.NewDecoder(resp.Body).Decode(&links); err != nil {
		return nil, err
	}
	return links, nil
}

func UnrestrictContainerLink(client *real_debrid.Client, link string) ([]string, error) {
	url := client.GetUrl("/unrestrict/containerLink")
	form := urlpkg.Values{}
	form.Add("link", link)
	req, err := http.NewRequest("POST", url.String(), strings.NewReader(form.Encode()))
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
	var links []string
	if err := json.NewDecoder(resp.Body).Decode(&links); err != nil {
		return nil, err
	}
	return links, nil
}
