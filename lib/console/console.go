package console

import (
	"fmt"

	"github.com/kinensake/gosearch/lib/core"
)

// Show syntax help.
func ShowHelp() {
	fmt.Println("Syntax: gosearch <package-keyword>")
	fmt.Println("Example: gosearch mongo")
}

// Intitialize interactive console display
//
// Parameters:
//	- keyword: package keyword
func Initialize(keyword string) {
	paths := core.GetPackagePath(keyword)
	if len(paths) == 0 {
		fmt.Println("No package found!!")
		return
	}
	fmt.Print("Top packages:\n\n")

	showChoices(paths)
	c := makeChoice(len(paths))
	core.ExecCommand(paths[c-1])
}

// Get package choice from user
//
// Parameters:
//	- optionLen: number of available options
// Return: package choice number
func makeChoice(optionLen int) int {
	var c int
	for {
		fmt.Print("\nChoose> ")
		fmt.Scanf("%d", &c)

		if c >= 1 && c <= optionLen {
			return c
		}
		fmt.Println("Invalid input, choose again!!!")
	}
}

// Display package choices
//
// Parameters:
//	- paths: package paths list
func showChoices(paths []string) {
	for i, path := range paths {
		fmt.Printf("%d. %s\n", i+1, path)
	}
}
