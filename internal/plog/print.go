package plog

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"io"
	"os"
	"strings"
)

const maxWidth = 100

var printer io.Writer = os.Stdout
var writer io.Writer = os.Stderr

func Print(a ...any) {
	_, _ = printer.Write([]byte(fmt.Sprint(a...)))
}

func Printf(format string, a ...any) {
	_, _ = printer.Write([]byte(fmt.Sprintf(format, a...)))
}

func Println(a ...any) {
	_, _ = printer.Write([]byte(fmt.Sprintln(a...)))
}

func Write(a ...any) {
	_, _ = writer.Write([]byte(fmt.Sprint(a...)))
}

func Writef(format string, a ...any) {
	_, _ = writer.Write([]byte(fmt.Sprintf(format, a...)))
}

func Writeln(a ...any) {
	_, _ = writer.Write([]byte(fmt.Sprintln(a...)))
}

func WriteTable(table [][]string) {
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

func Progress(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s\r", runewidth.Truncate(msg, maxWidth-3, "..."))
}
func ProgressEnd(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s\r", runewidth.FillLeft(" ", maxWidth))
	fmt.Println(msg)
}
