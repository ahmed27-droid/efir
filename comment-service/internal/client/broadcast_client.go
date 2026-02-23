package client

import (
	"comment-Service/internal/errs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BroadcastClient struct {
	baseURL string
}

func NewBroadcastClient(baseURL string) *BroadcastClient {
	return &BroadcastClient{baseURL: baseURL}
}

func (c *BroadcastClient) PostExists(postID uint) (bool, error) {
	url := fmt.Sprintf("%s/posts/%d", c.baseURL, postID)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("failed to close response body:", err)
		}
	}()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return false, errs.ErrBroadcastService
	}

	return true, nil
}

func (c *BroadcastClient) IsActive(postID uint) (bool, error) {
	url := fmt.Sprintf("%s/broadcasts/%d/status", c.baseURL, postID)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("failed to close response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return false, errs.ErrUnexpectedStatusCode
	}

	var result struct {
		Active bool `json:"active"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Active, nil
}
