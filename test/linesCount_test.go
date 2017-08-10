package test

import (
	"os"
	"WC/linesCount"
	"fmt"
	"testing"
)

func printResultLines(lines linesCount.List) {
	for i := 0; i < len(lines); i++ {
		fmt.Print(lines[i].GetText())
		fmt.Print(lines[i].GetLinesSize())
		fmt.Println(" ", lines[i].GetFilePath())
	}
}

func TestLinesCount(t *testing.T) {
	filePaths := os.Args[1:]

	if len(filePaths) == 0 {
		t.Error("Enter the filename(s)!")
	}

	lines := linesCount.LinesCount(filePaths)

	printResultLines(lines)
}
