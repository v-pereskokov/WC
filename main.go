package main

import (
	"os"
	"WC/linesCount"
	"fmt"
)

func printResultLines(lines linesCount.List) {
	for i := 0; i < len(lines); i++ {
		fmt.Print(lines[i].GetText())
		fmt.Print(lines[i].GetLinesSize())
		fmt.Println(" ", lines[i].GetFilePath())
	}
}

func main() {
	filePaths := os.Args[1:]

	if len(filePaths) == 0 {
		panic("Enter the filename(s)!")
	}

	//lines := linesCount.LinesCount(filePaths)

	//printResultLines(lines)
	linesCount.CountLines(filePaths)
}
