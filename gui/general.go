package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

const (
	viewFiles         = "files"
	viewLogs          = "logs"
	viewFileLogsPopup = "fileLogsPopup"
	viewFilter        = "filter"
	viewHelp          = "help"
	viewFilterPopup   = "filterPopup"

	vSeparator       = 30
	filterViewHeight = 3
	helpViewHeight   = 5
)

type view struct {
	draw   func(g *gocui.Gui) error
	events func(g *gocui.Gui) error
}

var views = map[string]view{
	"logs":   view{draw: logsDraw, events: logsEvents},
	"files":  view{draw: filesDraw, events: noEvents},
	"filter": view{draw: filterDraw, events: filterEvents},
	"help":   view{draw: helpDraw, events: helpEvents},
}

//InitializeGui initialize UI and lisplay filenames and logs
func InitializeGui(filenames []string) *gocui.Gui {
	loadContext(filenames, "")

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	g.Cursor = true
	g.SetManagerFunc(layout)

	for _, view := range views {
		if err := view.events(g); err != nil {
			log.Panicln(err)
		}
	}

	showFilenames(g, filenames)
	reloadLogs(g, nil)

	return g
}

func layout(g *gocui.Gui) error {
	for _, view := range views {
		if err := view.draw(g); err != nil {
			return err
		}
	}
	return nil
}

func excl(val int) int {
	return val - 1
}

func noEvents(g *gocui.Gui) error {
	return nil
}
