package main

import (
	"github.com/kingname/cgantt/task"
	"github.com/kingname/cgantt/utils"
	"fmt"
	"os"
	"time"
	"encoding/json"
	"io/ioutil"
)

func ParseJson(jsonStr string) []*task.ConfigedTask{
	t := []*task.ConfigedTask{}
	err := json.Unmarshal([]byte(jsonStr), &t)
	if nil != err{
		panic(err)
	}
	return t
}

func ReadAgenda() []*task.ConfigedTask {
	_, err := os.Stat("Agenda.json")
	if nil == err{
		jsonStr, err := ioutil.ReadFile("Agenda.json")
		if nil != err{
			panic(err)
		}
		return ParseJson(string(jsonStr))
	}
	return nil
}

func main() {
	width := utils.GetCurrentTerminalWidth()
	if width <= 0 {
		width = 100
	}
	fmt.Println("当前终端窗口宽度：", width)
	line1 := task.Line{"-", 10, "", 0, "=", 10}
	line2 := task.Line{"-", 11, "", 0, "=", 9}
	line3 := task.Line{"-", 8, "", 0, "=", 12}
	line4 := task.Line{"-", 3, "", 0, "=", 17}
	line5 := task.Line{"-", 2, "", 0, "=", 18}
	line6 := task.Line{"-", 7, "", 0, "=", 13}
	line7 := task.Line{"-", 7, "", 0, "=", 13}
	task1 := task.Task{"task1", line1, time.Date(2018,1,6, 0, 0, 0,0, time.UTC), time.Date(2018,1,20, 0, 0, 0, 0, time.UTC)}
	task2 := task.Task{"task2", line2, time.Date(2018,1,12, 0, 0, 0,0, time.UTC), time.Date(2018,1,16, 0, 0, 0, 0, time.UTC)}
	task3 := task.Task{"task3", line3,time.Date(2018,1,13, 0, 0, 0,0, time.UTC), time.Date(2018,1,18, 0, 0, 0, 0, time.UTC)}
	task4 := task.Task{"taskasdfasdfasdfsafasdf", line4,time.Date(2018,1,6, 0, 0, 0,0, time.UTC), time.Date(2018,1,17, 0, 0, 0, 0, time.UTC)}
	task5 := task.Task{"task5", line5,time.Date(2018,1,7, 0, 0, 0,0, time.UTC), time.Date(2018,1,9, 0, 0, 0, 0, time.UTC)}
	task6 := task.Task{"task6", line6,time.Date(2018,1,17, 0, 0, 0,0, time.UTC), time.Date(2018,1,21, 0, 0, 0, 0, time.UTC)}
	task7 := task.Task{"task7", line7,time.Date(2018,1,20, 0, 0, 0,0, time.UTC), time.Date(2018,1,21, 0, 0, 0, 0, time.UTC)}
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
