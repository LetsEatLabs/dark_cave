package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Locations object
type Location struct {
	Name               string          `json:"name"`
	CurrentDescription int             `json:"current_description"`
	Decriptions        []string        `json:"descriptions"`
	ConnectedLocations map[string]bool `json:"connected_locations"`
	Objects            []Object        `json:"objects"`
}

// Object object (lol)
// This is just for whats in a location. Not the object details. The details are
// Kept in objs.json just to keep things a little less cluttered.
type Object struct {
	Name       string `json:"name"`
	Pickup     bool   `json:"can_pickup"`
	Visible    bool   `json:"is_visible"`
	Interacted bool   `json:"interacted"`
	Total      int    `json:"total"`
	Details    string
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

func loadObjectDetails(g *Game) {
	var objs map[string]string

	content, err := ioutil.ReadFile("./objs.json")

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(content, &objs)

	for i := range g.locations {
		for t := range g.locations[i].Objects {
			g.locations[i].Objects[t].Details = objs[g.locations[i].Objects[t].Name]
		}
	}
}

// Reads the location description based on the current decription value in the Location
// object. Writes it to the terminal.
func readLocationDesc(loc string, g *Game) string {

	rtr := ""

	for l := range g.locations {
		if loc == g.locations[l].Name {
			desc_num := g.locations[l].CurrentDescription
			rtr = g.locations[l].Decriptions[desc_num]
		}
	}

	return rtr
}

func examineItem(g *Game, item []string) {

	itemName := strings.Join(item, "_")

	for i := range g.locations {
		if g.locations[i].Name == g.currentLocation {
			for t := range g.locations[i].Objects {
				if g.locations[i].Objects[t].Name == itemName {
					WriteOutputToTerminal(g, g.locations[i].Objects[t].Details)
					return
				}
			}
		}
	}

	// If the item is not found, complain
	WriteOutputToTerminal(g, fmt.Sprintf("You do not see %s here.", strings.Join(item, " ")))
	return
}

// Get a list of all items that are current in a location
// if visible=true then will only return items marked as visible
func getLocationItems(g *Game, loc string, visible bool) []string {
	var objNames []string

	for i := range g.locations {
		if g.locations[i].Name == loc {
			for t := range g.locations[i].Objects {
				ob := g.locations[i].Objects[t]
				obName := strings.Replace(ob.Name, "_", " ", -1)

				if visible {
					if ob.Visible {
						objNames = append(objNames, obName)
					}

				} else {
					objNames = append(objNames, obName)
				}

			}
		}
	}

	return objNames
}

// Get a list of all locations that are currented to the current location
// if visible=true then will only return items marked as visible
func getLocationConnectedLocations(g *Game, loc string, visible bool) []string {
	var locNames []string

	for i := range g.locations {
		if g.locations[i].Name == loc {
			for k, v := range g.locations[i].ConnectedLocations {
				kName := strings.Replace(k, "_", " ", -1)

				if visible {
					if v {
						locNames = append(locNames, kName)
					}
				} else {
					locNames = append(locNames, kName)
				}

			}
		}

	}

	return locNames
}
