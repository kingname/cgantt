package task

import (
	"time"
)

// Line represents the line, such: ----------==============
type Line struct {
	PreChar          string
	PreCharLength    int
	MiddleChar       string
	MiddleCharLength int
	PostChar         string
	PostCharLength   int
}

type Task struct {
	Name    string
	Gline   Line
	StartAt time.Time
	EndAt   time.Time
}

func DrawLine(line Line) string {
	fullLine := ""
	for i := 0; i < line.PreCharLength; i++ {
		fullLine += line.PreChar
	}
	for i := 0; i < line.MiddleCharLength; i++ {
		fullLine += line.MiddleChar
	}
	for i := 0; i < line.PostCharLength; i++ {
		fullLine += line.PostChar
	}
	return fullLine
}

//func DrawTask(task Task) {
//	return nil
//}
