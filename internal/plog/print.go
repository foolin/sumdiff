package plog

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"io"
	"os"
	"strings"
)

const maxWidth = 100

var Writer io.Writer = os.Stdout

func Print(a ...any) {
	_, _ = Writer.Write([]byte(fmt.Sprint(a...)))
}

func Printf(format string, a ...any) {
	_, _ = Writer.Write([]byte(fmt.Sprintf(format, a...)))
}

func Println(a ...any) {
	_, _ = Writer.Write([]byte(fmt.Sprintln(a...)))
}

func Progress(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s\r", runewidth.Truncate(msg, maxWidth-3, "..."))
}
func PrintEnd(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	Printf("%s\r", runewidth.FillLeft(" ", maxWidth))
	Println(msg)
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
