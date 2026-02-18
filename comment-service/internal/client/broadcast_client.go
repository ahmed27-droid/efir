package client

import (
	"commen-sService/internal/errors"
	"encoding/json"
	"fmt"
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
	defer resp.Body.Close()
	

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return false, errors.ErrBroadcastService
	}

	return true, nil
}

func (c *BroadcastClient) IsActive(postID uint) (bool, error) {
	url := fmt.Sprintf("%s/broadcasts/%d/status", c.baseURL, postID)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var result struct {
		Active bool `json:"active"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Active, nil
}

