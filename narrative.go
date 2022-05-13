package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Locations object
type Location struct {
	Name               string          `json:"name"`
	CurrentDescription int             `json:"current_description"`
	Decriptions        []string        `json:"descriptions"`
	ConnectedLocations map[string]bool `json:"connected_locations"`
}

func loadGameLocations(g *Game) {

	var decoded []Location
	content, err := ioutil.ReadFile("./loc.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &decoded)

	if err != nil {
		log.Fatal(err)
	}

	g.locations = decoded

}

// Reads the location description based on the current decription value in the Location
// object. Writes it to the terminal.
func readLocationDesc(loc string, g *Game) {
	for l := range g.locations {
		if loc == g.locations[l].Name {
			desc_num := g.locations[l].CurrentDescription
			WriteOutputToTerminal(g, g.locations[l].Decriptions[desc_num])
		}
	}
}
