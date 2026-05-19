package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/undetermiined/pokeCLI/internal/pokecache"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func getAndCache[T any](c *Client, url string) (T, error) {
	var zero T

	if val, ok := c.cache.Get(url); ok {
		var result T
		if err := json.Unmarshal(val, &result); err != nil {
			return zero, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return zero, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zero, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, err
	}

	var result T
	if err := json.Unmarshal(dat, &result); err != nil {
		return zero, err
	}
	c.cache.Add(url, dat)
	return result, nil
}
