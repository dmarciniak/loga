package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"loga/gui"
	"os"
)

const (
	textOneFilenameRequired = "At least one filename is required"
)

func main() {
	g := gui.InitializeGui(loadFilenames())
	defer g.Close()
	mainLoop(g)
}

func loadFilenames() []string {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println(textOneFilenameRequired)
		os.Exit(1)
	}
	return arguments[1:]
}

func mainLoop(g *gocui.Gui) {
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
