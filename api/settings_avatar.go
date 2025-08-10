package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	real_debrid "github.com/sushydev/real_debrid_go"
)

func UploadAvatarFile(client *real_debrid.Client, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	fw, err := w.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return err
	}
	if _, err := io.Copy(fw, f); err != nil {
		return err
	}
	w.Close()
	url := client.GetUrl("/settings/avatarFile")
	req, err := http.NewRequest("PUT", url.String(), &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return client.HandleResponseCode(resp, 204)
}

func DeleteAvatar(client *real_debrid.Client) error {
	url := client.GetUrl("/settings/avatarDelete")
	req, err := http.NewRequest("DELETE", url.String(), nil)
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
