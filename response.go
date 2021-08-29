package swapi

import (
	"encoding/json"
	"strings"
)

type apiResponse struct {
	Count int `json:"count"`
	Next string `json:"GetNext"`
	Previous string `json:"GetPrevious"`
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

	if _, err = r.do(req, r); err != nil {
		return r, err
	}

	return r, nil
}

func parseResult(r interface{}, out interface{}) error {
	jString, _ := json.Marshal(r)
	if err := json.Unmarshal(jString, out); err != nil {
		return err
	}
	return nil
}