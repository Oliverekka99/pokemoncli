package main

import (
	"fmt"

	"github.com/Oliverekka99/pokemoncli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	// startRepl()
}