package goverseerr

import "fmt"

type CacheStats struct {
	Hits   int `json:"hits"`
	Misses int `json:"misses"`
	Keys   int `json:"keys"`
	KSize  int `json:"ksize"`
	VSize  int `json:"vsize"`
}

type Cache struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Stats CacheStats `json:"stats"`
}

func (o *Overseerr) GetCacheStats() ([]*Cache, error) {
	var cache []*Cache
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&cache).Get("/settings/cache")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return cache, nil
}

func (o *Overseerr) FlushCache(cacheID string) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("cacheID", cacheID).
		Post("/settings/cache/{cacheID}/flush")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return fmt.Errorf("received non-204 status code (%d)", resp.StatusCode())
	}
	return nil
}
