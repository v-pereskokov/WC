package test

import (
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
	filePaths := []string{"input.txt", "top.txt"}

	lines := linesCount.LinesCount(filePaths)

	if lines[2].GetLinesSize() != 14 {
		t.Error("Error execute!")
	}

	printResultLines(lines)
}
