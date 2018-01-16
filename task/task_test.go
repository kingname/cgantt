package task

import (
	"testing"
)

func Test_Drawline(t *testing.T) {
	line := Line{"-", 10, "", 0, "=", 10}
	fullLine := DrawLine(line)
	answer := "----------=========="
	if answer == fullLine {
		t.Log("Pass")
	} else {
		t.Error("Draw Line Error.")
	}
}

func Test_DrawTaskArray(t *testing.T) {

	line1 := Line{"-", 10, "", 0, "=", 10}
	line4 := Line{"-", 3, "", 0, "=", 17}
	task1 := Task{"task1", line1}
	task4 := Task{"taskasdfasdfasdfsafasdf", line4}
	var taskArray []Task
	taskArray = append(taskArray, task1)
	taskArray = append(taskArray, task4)
	output := DrawTaskArray(taskArray)
	if output[0] == "|task1    | ----------==========|" && output[1] == "|taskasdf | ---=================|" {
		t.Log("Pass")
	} else {
		t.Error("绘图错误")
	}
}
