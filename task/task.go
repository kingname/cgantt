package task

import (
	// "time"
	"bytes"
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
	// StartAt time.Time
	// EndAt   time.Time
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

//DrawTask 将一个任务的任务名和线拼接在一起
func DrawTask(task Task, rjust int) string{
	taskName := task.Name
	var buffer bytes.Buffer
	if len(taskName) >= rjust {
		taskName = taskName[:rjust - 2]
	} else {
		for i:=0;i<rjust - len(taskName);i++{
			taskName += " "
		}
	}

	taskLine := DrawLine(task.Gline)
	buffer.WriteString("|")
	buffer.WriteString(taskName)
	buffer.WriteString(" | ")
	buffer.WriteString(taskLine)
	buffer.WriteString("|")
	return buffer.String()
}

//DrawTaskArray 把一系列任务拼接为长的字符串数组
func DrawTaskArray(taskArray []Task) []string{
	var taskLineArray []string
	for _, task := range taskArray{
		taskLineArray = append(taskLineArray, DrawTask(task, 10))
	}
	return taskLineArray
}
