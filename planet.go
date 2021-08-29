package swapi

import (
	"fmt"
)

// A Planet is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.
type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	ResidentURLs   []string `json:"residents"`
	FilmURLs       []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

// Planet retrieves the planet with the given id
func (c *Client) Planet(id int) (Planet, error) {
	req, err := c.newRequest(fmt.Sprintf("planets/%d", id))
	if err != nil {
		return Planet{}, err
	}

	var planet Planet

	if _, err = c.do(req, &planet); err != nil {
		return Planet{}, err
	}

	return planet, nil
}

// Planets retrieves all the planets as a paginated apiResponse
func (c *Client) Planets(page int) (*apiResponse, error) {
	response := apiResponse{
		Client: c,
	}

	path := fmt.Sprintf("planets/?page=%d", page)

	return response.get(path)
}

// AllPlanets retrieves all the planets as a paginated apiResponse
func (c *Client) AllPlanets() ([]Planet, error) {

	var planets []Planet
	resp, err := c.Planets(1)
	if err != nil {
		return []Planet{}, err
	}

	//planets = append(planets, parseResult(resp.results))
	fmt.Println(resp)
	return planets, err
}
