package main

import (
	"github.com/kingname/cgantt/task"
	"github.com/kingname/cgantt/utils"
)

func main() {
	width := utils.GetCurrentTerminalWidth()
	line1 := task.Line{"-", width / 2, "", 0, "=", width / 2}
	line2 := task.Line{"-", width / 3, "", 0, "=", 2 * width / 3}
	line3 := task.Line{"-", width / 4, "", 0, "=", 3 * width / 4}
	line4 := task.Line{"-", width / 5, "", 0, "=", 4 * width / 5}
	line5 := task.Line{"-", 5 * width / 6, "", 0, "=", width / 6}
	line6 := task.Line{"-", 3 * width / 7, "", 0, "=", 4 * width / 7}
	line7 := task.Line{"-", 4 * width / 8, "", 0, "=", width / 2}
	task.DrawLine(line1)
	task.DrawLine(line2)
	task.DrawLine(line3)
	task.DrawLine(line4)
	task.DrawLine(line5)
	task.DrawLine(line6)
	task.DrawLine(line7)
}
