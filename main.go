package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/kingname/cgantt/task"
)

func main() {
	width := getCurrentTerminalWidth()
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

func raiseErr(err error) {
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// getCurrentTerminalWidth do what its name represents it will do
func getCurrentTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	raiseErr(err)
	heightWidthArray := strings.Split(string(out), " ")
	width := heightWidthArray[1]
	widthInt, err := strconv.Atoi(strings.TrimSuffix(width, "\n"))
	raiseErr(err)
	return widthInt
}
