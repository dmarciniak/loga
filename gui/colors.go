package gui

import (
	"strconv"
)

var backgroundColors = [...]int{41, 42, 43, 44, 45, 46, 47}

func legendSmall(index int) string {
	return "\u001b[" + strconv.Itoa(backgroundColors[index+1%len(backgroundColors)]) + "m \x1b[0m "
}

func legend(index int) string {
	return "\u001b[" + strconv.Itoa(backgroundColors[index+1%len(backgroundColors)]) + "m  \x1b[0m "
}

func formatedFilter(str string) string {
	return "\u001b[32m" + str + "\u001b[0m"
}

func formatedDate(str string) string {
	return "\u001b[33m" + str + "\u001b[0m"
}

func formatedLogWithoutDate(str string) string {
	return "  \u001b[36m" + str + "\u001b[0m"
}

func formatedAlert(str string) string {
	return "\u001b[31m" + str + "\u001b[0m"
}
