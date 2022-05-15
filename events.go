package main

import "strings"

// This file is basically a giant list of scripted actions that can happen when 
// The player takes certain actions. Ex: Looking at X in Y room causes the description
// Of room Z to change

func checkForScripting(g *Game, cmd string, cmdArgs []string) {
	switch cmd {
	case "examine":
		handleExamineScripting(g, cmdArgs)
	case "goto":
		handleGoToScripting(g, cmdArgs)
	}
}

/////

// Updates the named location current_description value to the passed val
func updateLocationDescription(g *Game, loc string, val int) {
	for l := range g.locations {
			if loc == g.locations[l].Name {
				g.locations[l].CurrentDescription = val
				return
			}
		}
}

// Updates the named location name connected_location to the bool val passed
func updateConnectedLocationVisibility(g *Game, loc string, connected_loc string, val bool) {
	for l := range g.locations {
			if loc == g.locations[l].Name {
				for k, _ := range g.locations[l].ConnectedLocations {
					if k == connected_loc {
						g.locations[l].ConnectedLocations[connected_loc] = val
						return
					}
					
				}
				
			}
		}
}

func updateLocationItemVisibility(g *Game, loc string, item string, val bool) {
	for l := range g.locations {
			if loc == g.locations[l].Name {
				for i := range g.locations[l].Objects {
					if g.locations[l].Objects[i].Name == item {
						g.locations[l].Objects[i].Visible = val
						return
					}
					
				}
				
			}
		}
}

/////


// If something has scripting associated with examining it, put it in here
func handleExamineScripting(g *Game, cmdArgs []string) {
	exobj := strings.Join(cmdArgs, "_")

	// Picking up the photo in the parlor
	if exobj == "photo" && g.currentLocation == "parlor" {
		updateLocationDescription(g, "hallway", 1)
		updateLocationDescription(g, "kitchen", 1)
		updateConnectedLocationVisibility(g, "kitchen", "pantry", true)
	}
}



// If a place has scripting associated with going to it, put it in here
func handleGoToScripting(g *Game, cmdArgs []string) {
	exobj := strings.Join(cmdArgs, "_")

	// Entering the kitchen
	if exobj == "kitchen" {
		updateLocationDescription(g, "garden", 1)
		updateLocationItemVisibility(g, "garden", "tomato_seeds", false)
	}

	if exobj == "pantry" {
		updateLocationItemVisibility(g, "kitchen", "recipe", false)
	}
}