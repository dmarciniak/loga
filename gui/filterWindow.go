package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	textFilterKeys = " ctrl + f to set"
)

func filterDraw(g *gocui.Gui) error {
	_, maxY := g.Size()
	if v, err := g.SetView(viewFilter, 0, maxY-(helpViewHeight+filterViewHeight), excl(vSeparator), excl(maxY-helpViewHeight)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = viewFilter
		v.Editable = false
		v.Wrap = false
		fmt.Fprintln(v, textFilterKeys)
	}
	return nil
}

func filterEvents(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlF, gocui.ModNone, showFilterPopup); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewFilterPopup, gocui.KeyEnter, gocui.ModNone, closeFilterPopup); err != nil {
		return err
	}
	return nil
}

func showFilterPopup(g *gocui.Gui, _ *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(viewFilterPopup, maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Set filter (regex)"
		v.Editable = true
		fmt.Fprintln(v, ctx.filter)
		if _, err := g.SetCurrentView(viewFilterPopup); err != nil {
			return err
		}
	}
	return nil
}

func closeFilterPopup(g *gocui.Gui, v *gocui.View) error {
	filter, _ := v.Line(0)
	updateContext(filter)

	if err := g.DeleteView(viewFilterPopup); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(viewLogs); err != nil {
		return err
	}
	updateFilter(g)
	reloadLogs(g, nil)
	return nil
}

func updateFilter(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewFilter)
		if err != nil {
			return err
		}
		v.Clear()
		if ctx.filter == "" {
			fmt.Fprintln(v, textFilterKeys)
		} else {
			fmt.Fprintln(v, formatedFilter(ctx.filter))
		}
		return nil
	})
}

func invalidRegex(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewFilter)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, formatedAlert("Invalid regex"))
		return nil
	})
}
