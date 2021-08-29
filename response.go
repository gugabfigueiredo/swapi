package swapi

import "strings"

type apiResponse struct {
	count int `json:"count"`
	next string `json:"next"`
	previous string `json:"previous"`
	results []interface{} `json:"result"`
	*Client
}

func (r *apiResponse) HasNext() bool {
	return r.next != ""
}

func (r *apiResponse) HasPrevious() bool {
	return r.next != ""
}

func (r *apiResponse) get(path string) (*apiResponse, error) {
	endpoint := strings.Split(path, r.basePath)[1]
	req, err := r.newRequest(endpoint)
	if err != nil {
		return r, err
	}

	var resp apiResponse
	if _, err = r.do(req, &resp); err != nil {
		return r, err
	}
	resp.Client = r.Client

	return &resp, nil
}

func (r *apiResponse) Next() (*apiResponse, error) {
	return r.get(r.next)
}

func (r *apiResponse) Previous() (*apiResponse, error) {
	return r.get(r.previous)
}