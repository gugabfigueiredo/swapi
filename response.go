package swapi

import (
	"bytes"
	"encoding/json"
	"strings"
)

type apiResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []interface{} `json:"results"`
	*Client `json:"client,omitempty"`
}

func (r *apiResponse) HasNext() bool {
	return r.Next != ""
}

func (r *apiResponse) HasPrevious() bool {
	return r.Previous != ""
}

func (r *apiResponse) GetNext() (*apiResponse, error) {
	return r.get(r.Next)
}

func (r *apiResponse) GetPrevious() (*apiResponse, error) {
	return r.get(r.Previous)
}

func (r *apiResponse) get(path string) (*apiResponse, error) {
	splitPath := strings.Split(path, r.basePath)
	endpoint := splitPath[len(splitPath)-1]
	req, err := r.newRequest(endpoint)
	if err != nil {
		return r, err
	}

	resp := &apiResponse{
		Client: r.Client,
	}
	if _, err = r.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func parseResult(r interface{}, out interface{}) error {
	jString, _ := json.Marshal(r)
	jBytes := bytes.NewReader(jString)
	if err := json.NewDecoder(jBytes).Decode(out); err != nil {
		return err
	}
	return nil
}