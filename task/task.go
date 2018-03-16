package task

import (
	"os"
	"time"
	"fmt"
	"bytes"
	"github.com/kingname/cgantt/utils"
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

type ConfigedTask struct {
	Name string
	StartAt time.Time
	EndAt time.Time
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

	task.Gline = CalcLine(time.Now(), task)
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

func clacWidthPerDay() int{
	terminalWidth := utils.GetCurrentTerminalWidth()
	if terminalWidth <= 20 {
		fmt.Println("终端窗口太小！不能绘制甘特图")
		os.Exit(3)
	}
	widthPerDay := (terminalWidth - 20) / 15
	return widthPerDay
}

func CalcLine(now time.Time, task Task) Line{
	startTime := task.StartAt
	endTime := task.EndAt
	dayBefore7days := startTime.Add(-24 * time.Hour * 7)
	dayAfter7days := startTime.Add(24 * time.Hour * 7)
	if startTime.Before(dayBefore7days) {
		startTime = dayBefore7days
	}
	if endTime.After(dayBefore7days){
		endTime = dayAfter7days
	}
	preCharNum := int(startTime.Sub(dayBefore7days).Hours() / 24)
	afterCharNum := int(dayAfter7days.Sub(endTime).Hours() / 24)
	middleCharNum := int(endTime.Sub(startTime).Hours() / 24)
	line := Line{" ", preCharNum, "=", middleCharNum, " ", afterCharNum}
	return line
}