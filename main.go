package main

import (
	"os"

	"github.com/kinensake/gosearch/lib/console"
)

func main() {
	if len(os.Args) != 2 {
		console.ShowHelp()
		return
	}
	keyword := os.Args[1]
	console.Initialize(keyword)
}
