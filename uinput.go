package main

import (
	"fmt"
	"os"
	"strings"
)

// Definitions

var knownCommands = []string{
	"goto",
	"help",
	"look",
	"examine",
	"exit",
}

var commandUsage = map[string]string{
	"goto":    "Go to a location from your current location, if possible",
	"help":    "Displays this information",
	"look":    "You look around your current surroundings",
	"examine": "Investigate a particular item in the area",
	"exit":    "You go home",
}

// Writes the user input to the terminal and then clears the user input field
func MoveInputToTerminal(g *Game) {
	// First clear the screen so it is fresh every time the player hits enter
	g.text = ""

	g.text += fmt.Sprintf("You: %s\n", strings.Replace(g.playerInputText, "> ", "", 1))
	g.playerInputText = g.ps1
}

// Writes a string exactly as passed to the terminal, accounting for column width
func WriteOutputToTerminal(g *Game, str string) {

	newText := ""

	a := 0
	for i := range str {

		letter := string(str[i])
		newText += letter

		// Restart our column counter if we have a new line
		if letter == "\n" {
			a = 0
		}

		if a%textWidth == 0 && a != 0 {
			newText += "\n"
			a = 0
			continue
		}

		// Check if there is a space coming up soon and break early to prevent
		// Word wrap if we can

		spaceLeft := textWidth - a - 1 // subtract one extra to ensure we do not go over

		if spaceLeft < wrapDistance {
			if spaceLeft > 0 {
				if letter == " " {
					for t := 0; t < spaceLeft; t++ {
						if i+t < len(str) {
							if string(str[t+i]) == " " {
								newText += "\n"
								a = 0
								break
							}
						}

					}
				}

			}
		}

		a = a + 1
	}

	g.text += fmt.Sprintf("%s", newText)

}

// Splits the command on space, and makes all lower case
func ParseUserCommand(command string) []string {
	splitCommand := strings.Split(command, " ")
	var littleCommand []string

	for t := range splitCommand {

		// Skip the PS1
		if t == 0 {
			continue
		}

		littleCommand = append(littleCommand, strings.ToLower(splitCommand[t]))
	}

	return littleCommand
}

// Checks if a string is in a list of strings
func CheckIfListHasString(str string, arr []string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}
	}
	return false
}

// Determines what command was issues and acts as a helper function to
// Call that command's methods
func HandleCommand(g *Game, command []string) {

	if command[0] == "debug" {
		if command[1] == "on" {
			g.isDebug = true
			WriteOutputToTerminal(g, "Debug mode is enabled.")
			return
		}

		if command[1] == "off" {
			g.isDebug = false
			WriteOutputToTerminal(g, "Debug mode is disabled.")
			return
		}

	}

	// Do we know this command?
	if !CheckIfListHasString(command[0], knownCommands) {
		writeStr := fmt.Sprintf("I do not know what %s means. I know these things: \n- %s",
			command[0],
			strings.Join(knownCommands, "\n- "))

		WriteOutputToTerminal(g, writeStr)

		return
	}

	// Are we quitting?
	if command[0] == "exit" {
		os.Exit(0)
	}

	// Do you need help?
	if command[0] == "help" {
		writeStr := "Here is what you can currently do:\n"
		for k, v := range commandUsage {
			writeStr += fmt.Sprintf("-%s: %s\n", k, v)
		}

		WriteOutputToTerminal(g, writeStr)
		return

	}

	// If we look, read the current description of the Location we are in
	if command[0] == "look" {
		writeStr := getFullAreaDescription(g)
		WriteOutputToTerminal(g, writeStr)
		checkForScripting(g, command[0], command[1:])

	}

	// Examine an item
	if command[0] == "examine" {
		examineItem(g, command[1:])
		checkForScripting(g, command[0], command[1:])

	}

	// Go somewhere (which also does a 'look' in the new area)
	if command[0] == "goto" {
		realLoc := goToLocation(g, command[1:])

		if realLoc {
			writeStr := getFullAreaDescription(g)
			WriteOutputToTerminal(g, writeStr)
		}

		// If we are in debug mode, check for scripting even if we cannot
		// Access this area now, because screw you its debug mode.
		if g.isDebug {
			checkForScripting(g, command[0], command[1:])
		}

	}
}

// Returns a full description of the current area that a player is standing in.
func getFullAreaDescription(g *Game) string {
	locDesc := readLocationDesc(g.currentLocation, g)

	// List of items that the player can interact with
	locObjs := getLocationItems(g, g.currentLocation, true)

	// List of places the player can currently go
	conLocs := getLocationConnectedLocations(g, g.currentLocation, true)

	// Combine the items
	writeStr := fmt.Sprintf("%s\n\nYou can 'examine' these items: %s\n",
		locDesc,
		strings.Join(locObjs, ", "))

	writeStr += fmt.Sprintf("You can 'goto' these places: %s\n", strings.Join(conLocs, ", "))

	return writeStr
}
