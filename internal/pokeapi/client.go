package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/abdullahsamir74/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Client talks to PokéAPI and caches successful responses in memory.
type Client struct {
	cache      *pokecache.Cache
	httpClient *http.Client
}

// NewClient returns a client whose responses are cached for cacheInterval.
func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// get fetches url and decodes the JSON body into target, caching the raw body
// on success. Cached responses are returned without hitting the network.
func (c *Client) get(url string, target any) error {
	if data, ok := c.cache.Get(url); ok {
		fmt.Printf("Cache hit for %s\n", url)
		return json.Unmarshal(data, target)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: unexpected status %s", url, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, target); err != nil {
		return err
	}

	c.cache.Add(url, data)
	return nil
}
