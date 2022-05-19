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
			for k := range g.locations[l].ConnectedLocations {
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

func updateLocationItemDescription(g *Game, loc string, item string, val string) {
	for l := range g.locations {
		if loc == g.locations[l].Name {
			for i := range g.locations[l].Objects {
				if g.locations[l].Objects[i].Name == item {
					g.locations[l].Objects[i].Details = val
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
		updateLocationItemVisibility(g, "hallway", "family_photo", true)
	}

	// Picking up the keys in the pantry
	if exobj == "keys" && g.currentLocation == "pantry" {
		updateLocationDescription(g, "parlor", 1)
		updateConnectedLocationVisibility(g, "parlor", "courtyard", true)
		updateLocationItemVisibility(g, "pantry", "keys", false)
	}

	// Picking up the journal in the quarters
	if exobj == "journal" && g.currentLocation == "quarters" {
		updateLocationItemVisibility(g, "quarters", "glasses", true)
		updateLocationItemDescription(g, "cave_entrance", "blue_book", "'The time has come. We have waited patiently for over a decade. We have suffered through his annoying happiness and jarring wholesome attitude. Our spawn is now 10 years old - old enough for the ritual to take place. 10 years I have dealt with this annoying man, and now I shall transfer his mind into our spawn to create a genius. If everything works then he should remain here in a vegetative state forever. The only one who can undo this spell is the spawn itself on its own free will, awakening him.'")
	}

	// Picking up the glasses in the quarters
	if exobj == "glasses" && g.currentLocation == "quarters" {
		updateLocationItemDescription(g, "cave_entrance", "blue_book", "'The time has come. We have waited patiently for over a decade. We have suffered through his annoying happiness and jarring wholesome attitude. Our spawn is now 10 years old - old enough for the ritual to take place. 10 years I have dealt with this annoying man, and now I shall transfer his mind into our spawn to create a genius. If everything works then he should remain here in a vegetative state forever. The only one who can undo this spell is the spawn itself on its own free will, awakening him.'")
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
		updateLocationItemVisibility(g, "kitchen", "recipe", true)
	}
}
