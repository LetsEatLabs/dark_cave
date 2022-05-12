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

// Writes the user input to the terminal and then clears the user input field
func MoveInputToTerminal(g *Game) {
	g.text += fmt.Sprintf("You: %s\n", strings.Replace(g.playerInputText, "> ", "", 1))
	g.playerInputText = g.ps1
}

// Writes a string exactly as passed to the terminal, accounting for column width
func WriteOutputToTerminal(g *Game, str string) {
	g.text += fmt.Sprintf("%s\n", str)

	a := 0

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

	// If we look, read the current description of the Location we are in
	if command[0] == "look" {
		readLocationDesc(g.currentLocation, g)
	}
}

func RepeatKeyPressed(key ebiten.Key) {
	// pass
}
