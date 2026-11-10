package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Printf("Cache hit for %s\n", fullURL)
		var resp LocationAreasResp
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return resp, err
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreasResp{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	var locationAreasResp LocationAreasResp
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)
	return locationAreasResp, nil

}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := fmt.Sprintf("/location-area/%s", locationAreaName)
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Printf("Cache hit for %s\n", fullURL)
		var resp LocationArea
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return LocationArea{}, err
		}
		return resp, err
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	var locationAreasResp LocationArea
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)
	return locationAreasResp, nil

}
