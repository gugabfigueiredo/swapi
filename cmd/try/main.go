package main

import (
	"fmt"

	"github.com/gugabfigueiredo/swapi"
)

func main() {
	c := swapi.DefaultClient

	if planets, err := c.AllPlanets(); err == nil {
		fmt.Println("planets : ", planets)
	}


}