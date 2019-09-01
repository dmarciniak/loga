package gui

import (
	"fmt"
	"strings"
	"github.com/jroimartin/gocui"
	"github.com/dmarciniak/loge"
)

func showFileLogs(g *gocui.Gui, line lineInfo) error {
	maxX, maxY := g.Size()
	filename := ctx.filenames[line.fileIndex]
	if v, err := g.SetView(viewFileLogsPopup, 4, 3, maxX-4, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = filename + " (ctrl + c to hide popup)"
		v.Editable = false
		v.Wrap = false
		if _, err := g.SetCurrentView(viewFileLogsPopup); err != nil {
			return err
		}

		fileOutput := loge.LogLoader([]string{filename})
		for entry := <-fileOutput; !entry.IsEOF(); entry = <-fileOutput {
			currentLog := entry.Log

			if entry.LineNo == line.fileLineNumber {
				currentLog = formatedCurrentLine(currentLog)
			} else if entry.IsEmptyDate() {
				currentLog = formatedLogWithoutDate(currentLog)
			} else {
				currentLog = strings.Replace(currentLog, entry.RawDate, formatedDate(entry.RawDate), -1)
			}
			
			fmt.Fprintln(v, currentLog)
		}
		v.SetOrigin(0, line.fileLineNumber)
	}
	return nil
}

func logsPopupEvents(g *gocui.Gui) error {
	if err := g.SetKeybinding(viewFileLogsPopup, gocui.KeyCtrlC, gocui.ModNone, closeFileLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewFileLogsPopup, gocui.KeyArrowUp, gocui.ModNone, scrollUpLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewFileLogsPopup, gocui.KeyArrowDown, gocui.ModNone, scrollDownLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewFileLogsPopup, gocui.KeyArrowLeft, gocui.ModNone, scrollLeftLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewFileLogsPopup, gocui.KeyArrowRight, gocui.ModNone, scrollRightLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewFileLogsPopup, gocui.KeyPgdn, gocui.ModNone, scrollPgdnLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewFileLogsPopup, gocui.KeyPgup, gocui.ModNone, scrollPgupLogs); err != nil {
		return err
	}
	return nil
}

func closeFileLogs(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView(viewFileLogsPopup); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(viewLogs); err != nil {
		return err
	}
	return nil
}


