package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	textShowFile  = " enter to show whole file"
	textResetKeys = " ctrl + r to reset"
	textQuitKeys  = " ctrl + c to quit"
)

func helpDraw(g *gocui.Gui) error {
	_, maxY := g.Size()
	if v, err := g.SetView(viewHelp, 0, maxY-helpViewHeight, excl(vSeparator), excl(maxY)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = viewHelp
		v.Editable = false
		v.Wrap = false
		fmt.Fprintln(v, textShowFile)
		fmt.Fprintln(v, textResetKeys)
		fmt.Fprintln(v, textQuitKeys)
	}
	return nil
}

func helpEvents(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlR, gocui.ModNone, reset); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewLogs, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func reset(g *gocui.Gui, v *gocui.View) error {
	updateContext("")
	updateFilter(g)
	return reloadLogs(g, v)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
