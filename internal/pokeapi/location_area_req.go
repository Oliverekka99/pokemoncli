package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (*LocationAreaResp, error) {
	endpoint := "location-area"
	fullURL := baseURL + endpoint
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return &LocationAreaResp{}, err
	}
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return &LocationAreaResp{}, err
	}
	if resp.StatusCode > 399 {
		return &LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return &LocationAreaResp{}, fmt.Errorf("error unmarshalling response: %v", err)
	}


	defer resp.Body.Close()
	return &locationAreaResp, nil
}