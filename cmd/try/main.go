package main

import (
	"fmt"

	"github.com/gugabfigueiredo/swapi"
)

func main() {
	c := swapi.DefaultClient

	if atst, err := c.Vehicle(19); err == nil {
		fmt.Println("name: ", atst.Name)
		fmt.Println("model:", atst.Model)
	}
}