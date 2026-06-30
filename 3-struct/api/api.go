package api

import (
	"3-struct/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseURL = "https://api.jsonbin.io/v3"

type Client struct {
	apiKey string
}

func NewClient(cfg *config.Config) *Client {
	return &Client{apiKey: cfg.Key}
}

func (c *Client) setHeaders(req *http.Request, withBody bool) {
	req.Header.Set("X-Master-Key", c.apiKey)
	if withBody {
		req.Header.Set("Content-Type", "application/json")

	}
}

type createResponse struct {
	Metadata struct {
		ID string `json:"id"`
	} `json:"metadata"`
}
type getResponse struct {
	Record json.RawMessage `json:"record"`
}

func (c *Client) CreateBin(data []byte, name string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, baseURL+"/b", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	c.setHeaders(req, true)
	req.Header.Set("X-Bin-Name", name)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api error %d: %s", resp.StatusCode, body)
	}
	var result createResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result.Metadata.ID, nil
}

func (c *Client) GetBin(id string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+"/b/"+id, nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req, false)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error %d: %s", resp.StatusCode, body)
	}
	var result getResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Record, nil
}

func (c *Client) UpdateBin(id string, data []byte) error {
	req, err := http.NewRequest(http.MethodPut, baseURL+"/b/"+id, bytes.NewReader(data))
	if err != nil {
		return err
	}
	c.setHeaders(req, true)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("api error %d: %s", resp.StatusCode, body)
	}
	return nil
}
func (c *Client) DeleteBin(id string) error {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/b/"+id, nil)
	if err != nil {
		return err
	}
	c.setHeaders(req, false)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("api error %d: %s", resp.StatusCode, body)
	}
	return nil
}
