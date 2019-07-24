package gui

type context struct {
	filenames []string
	filter    string
}

var ctx context

func loadContext(filenames []string, filter string) {
	ctx = context{filenames: filenames, filter: filter}
}

func updateContext(filter string) {
	ctx = context{filenames: ctx.filenames, filter: filter}
}
