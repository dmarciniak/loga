package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func filesDraw(g *gocui.Gui) error {
	_, maxY := g.Size()
	if v, err := g.SetView(viewFiles, 0, 0, excl(vSeparator), excl(maxY-(helpViewHeight+filterViewHeight))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = viewFiles
		v.Editable = false
		v.Wrap = false
	}
	return nil
}

func showFilenames(g *gocui.Gui, filenames []string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewFiles)
		if err != nil {
			return err
		}
		for index, filename := range filenames {
			fmt.Fprintln(v, legend(index)+filename)
		}
		return nil
	})
}
