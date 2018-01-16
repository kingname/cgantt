package main

import (
	"github.com/kingname/cgantt/task"
	"github.com/kingname/cgantt/utils"
	"fmt"
)

func main() {
	width := utils.GetCurrentTerminalWidth()
	fmt.Println("当前终端窗口宽度：", width)
	line1 := task.Line{"-", 10, "", 0, "=", 10}
	line2 := task.Line{"-", 11, "", 0, "=", 9}
	line3 := task.Line{"-", 8, "", 0, "=", 12}
	line4 := task.Line{"-", 3, "", 0, "=", 17}
	line5 := task.Line{"-", 2, "", 0, "=", 18}
	line6 := task.Line{"-", 7, "", 0, "=", 13}
	line7 := task.Line{"-", 18, "", 0, "=", 2}
	task1 := task.Task{"task1", line1}
	task2 := task.Task{"task2", line2}
	task3 := task.Task{"task3", line3}
	task4 := task.Task{"taskasdfasdfasdfsafasdf", line4}
	task5 := task.Task{"task5", line5}
	task6 := task.Task{"task6", line6}
	task7 := task.Task{"task7", line7}
	var taskArray []task.Task
	taskArray = append(taskArray, task1)
	taskArray = append(taskArray, task2)
	taskArray = append(taskArray, task3)
	taskArray = append(taskArray, task4)
	taskArray = append(taskArray, task5)
	taskArray = append(taskArray, task6)
	taskArray = append(taskArray, task7)
	output := task.DrawTaskArray(taskArray)
	fmt.Println(task.DrawLine(task.Line{"*", 33, "", 0, "", 0}))
	for _, line := range output{
		fmt.Println(line)
	}
	fmt.Println(task.DrawLine(task.Line{"*", 33, "", 0, "", 0}))
}
