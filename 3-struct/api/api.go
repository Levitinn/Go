package api

import "3-struct/config"

type Client struct {
	apiKey string
}

func NewClient(cfg *config.Config) *Client {
	return &Client{apiKey: cfg.Key}
}

func (c *Client) APIKey() string {
	return c.apiKey
}
