package plog

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"strings"
)

func Print(message string) {
	fmt.Println(message)
}

func PrintProgress(message string) {
	width := 100
	fmt.Printf("%s\r", runewidth.Wrap(runewidth.Truncate(message, width-97, "..."), width))
}
func PrintProgressEnd(message string) {
	width := 100
	fmt.Printf("%s\r\n", runewidth.Wrap(runewidth.Truncate(message, width-97, "..."), width))
}

func PrintError(path string, message string) {
	fmt.Printf("| %-60s | %s |\n", runewidth.Truncate(path, 50, "..."), message)
	//fmt.Printf("| %-20s | %-20s |\n", path, message)
}

func PrintTable(table [][]string) {
	// get number of columns from the first table row
	columnLengths := make([]int, len(table[0]))
	for _, line := range table {
		for i, val := range line {
			if len(val) > columnLengths[i] {
				columnLengths[i] = len(val)
			}
		}
	}

	var lineLength int
	for _, c := range columnLengths {
		lineLength += c + 3 // +3 for 3 additional characters before and after each field: "| %s "
	}
	lineLength += 1 // +1 for the last "|" in the line

	for i, line := range table {
		if i == 0 { // table header
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
		for j, val := range line {
			fmt.Printf("| %-*s ", columnLengths[j], val)
			if j == len(line)-1 {
				fmt.Printf("|\n")
			}
		}
		if i == 0 || i == len(table)-1 { // table header or last line
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
	}
}
