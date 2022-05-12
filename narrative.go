package main

import (
	"encoding/json"
	"io/ioutil"
)

// Locations object
type Locations struct {
	Places []struct {
		Location struct {
			Name               string   `json:"name"`
			CurrentDescription int      `json:"current_description"`
			Decriptions        []string `json:"descriptions"`
		}
	}
}

func loadGameLocations(g *Game) {

	var Loc Locations
	content, err := ioutil.ReadFile("./loc.json")

	err = json.Unmarshal(content, &Loc)

	if err != nil {
		print(err)
	}

	g.locations = Loc

}

func readLocationDesc(loc string, g *Game) {

}
