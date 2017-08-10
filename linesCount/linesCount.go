package linesCount

import (
	"io/ioutil"
	"bytes"
)

type ResultCount struct {
	text      string
	linesSize int
	filePath  string
}

type List []ResultCount

func (l *List) append(textValue string, linesSizeValue int, fileNameValue string) {
	e := ResultCount{
		text:      textValue,
		linesSize: linesSizeValue,
		filePath:  fileNameValue,
	}

	*l = append(*l, e)
}

func (r *ResultCount) GetText() string {
	return r.text
}

func (r *ResultCount) GetLinesSize() int {
	return r.linesSize
}

func (r *ResultCount) GetFilePath() string {
	return r.filePath
}

func getLine(file []byte) int {
	return bytes.Count(file, []byte{'\n'})
}

func LinesCount(filePaths []string) List {
	total := 0
	result := List{}

	for _, value := range filePaths {
		file, error := ioutil.ReadFile(value)

		if error != nil {
			panic("No such file or directory!")
		}

		currentLength := getLine(file)
		total += currentLength

		result.append("      ", currentLength, value)
	}

	if len(filePaths) > 1 {
		result.append("      ", total, "total")
	}

	return result
}
