package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// Definitions

var knownCommands = []string{
	"help",
	"look",
	"exit",
}

var commandUsage = map[string]string {
	"help": "Displays this information",
	"look": "You look around your current surroundings",
	"exit": "You go home",
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

		a = a+1
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

	}

	// If we look, read the current description of the Location we are in
	if command[0] == "look" {
		readLocationDesc(g.currentLocation, g)
	}
}

func RepeatKeyPressed(key ebiten.Key) {
	// pass
}
