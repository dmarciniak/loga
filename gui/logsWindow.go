package gui

import (
	"strconv"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/dmarciniak/loge"
	"regexp"
	"strings"
	"sync"
)

const (
	logsLimit = 1000
)

var (
	isAllLogLoaded = false
	logsLoadingMutex sync.Mutex
)

func logsDraw(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(viewLogs, vSeparator, 0, excl(maxX), excl(maxY)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = viewLogs
		v.Editable = false
		v.Wrap = false
		if _, err := g.SetCurrentView(viewLogs); err != nil {
			return err
		}
	}
	return nil
}

func logsEvents(g *gocui.Gui) error {
	if err := g.SetKeybinding(viewLogs, gocui.KeyArrowUp, gocui.ModNone, scrollUpLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewLogs, gocui.KeyArrowDown, gocui.ModNone, scrollDownLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewLogs, gocui.KeyArrowLeft, gocui.ModNone, scrollLeftLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewLogs, gocui.KeyArrowRight, gocui.ModNone, scrollRightLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewLogs, gocui.KeyPgdn, gocui.ModNone, scrollPgdnLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding(viewLogs, gocui.KeyPgup, gocui.ModNone, scrollPgupLogs); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlN, gocui.ModNone, loadNextLogs); err != nil {
		return err
	}
	return nil
}

func scrollUpLogs(g *gocui.Gui, _ *gocui.View) error {
	moveLogsScreen(g, 0, -1)
	return nil
}

func scrollDownLogs(g *gocui.Gui, _ *gocui.View) error {
	moveLogsScreen(g, 0, 1)
	return nil
}

func scrollPgupLogs(g *gocui.Gui, _ *gocui.View) error {
	_, maxY := g.Size()
	moveLogsScreen(g, 0, -(maxY - 2))
	return nil
}

func scrollPgdnLogs(g *gocui.Gui, _ *gocui.View) error {
	_, maxY := g.Size()
	moveLogsScreen(g, 0, maxY-2)
	return nil
}

func scrollLeftLogs(g *gocui.Gui, _ *gocui.View) error {
	moveLogsScreen(g, -1, 0)
	return nil
}

func scrollRightLogs(g *gocui.Gui, _ *gocui.View) error {
	moveLogsScreen(g, 1, 0)
	return nil
}

func loadNextLogs(g *gocui.Gui, _ *gocui.View) error {
	writeLogs(g, ctx.output)
	return nil
}

func resetLogsScreen(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewLogs)
		if err != nil {
			return err
		}
		v.SetOrigin(0, 0)
		v.Clear()
		isAllLogLoaded = false
		return nil
	})
}

func moveLogsScreen(g *gocui.Gui, moveX, moveY int) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewLogs)
		if err != nil {
			return err
		}

		x, y := v.Origin()
		v.SetOrigin(x+moveX, y+moveY)
		return nil
	})
}

func reloadLogs(g *gocui.Gui, _ *gocui.View) error {
	ctx.output = loge.LogLoader(ctx.filenames)
	resetLogsScreen(g)
	writeLogs(g, ctx.output)
	return nil
}

func writeLogs(g *gocui.Gui, output <-chan loge.LogEntry) {
	g.Update(func(g *gocui.Gui) error {

		logsLoadingMutex.Lock()
		defer logsLoadingMutex.Unlock()

		v, err := g.View(viewLogs)
		if err != nil {
			return err
		}

		if isAllLogLoaded {
			fmt.Fprintln(v, formatedAlert("All logs loaded"));
			return nil
		}

		var rgx *regexp.Regexp
		if ctx.filter != "" {
			var err error
			if rgx, err = regexp.Compile(ctx.filter); err != nil {
				invalidRegex(g)
				return nil
			}
		}

		for i := 0; i <= logsLimit; {
			entry := <-output

			if entry.IsEOF() {
				isAllLogLoaded = true
				return nil
			}

			formatedLog := entry.Log

			findedString := ""
			if rgx != nil {
				findedString = rgx.FindString(entry.Log)
				if findedString == "" {
					continue
				}
			}

			if rgx != nil && findedString != "" {
				formatedLog = strings.Replace(formatedLog, findedString, formatedFilter(findedString), -1)
			}

			if entry.IsEmptyDate() {
				fmt.Fprintln(v, legend(entry.FileID)+formatedLogWithoutDate(formatedLog))
			} else {
				formatedLog = strings.Replace(formatedLog, entry.RawDate, formatedDate(entry.RawDate), -1)
				fmt.Fprintln(v, legend(entry.FileID)+formatedLog)
			}
			i++
		}

		if !isAllLogLoaded {
			fmt.Fprintln(v, formatedAlert("Loaded " + strconv.Itoa(logsLimit) + " logs. Press ctrl + n to load next logs"));
		}

		return nil
	})
}
