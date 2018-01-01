package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	width := getCurrentTerminalWidth()
	line1 := Line{"-", width / 2, "", 0, "=", width / 2}
	line2 := Line{"-", width / 3, "", 0, "=", 2 * width / 3}
	line3 := Line{"-", width / 4, "", 0, "=", 3 * width / 4}
	line4 := Line{"-", width / 5, "", 0, "=", 4 * width / 5}
	line5 := Line{"-", 5 * width / 6, "", 0, "=", width / 6}
	line6 := Line{"-", 3 * width / 7, "", 0, "=", 4 * width / 7}
	line7 := Line{"-", 4 * width / 8, "", 0, "=", width / 2}
	DrawLine(line1)
	DrawLine(line2)
	DrawLine(line3)
	DrawLine(line4)
	DrawLine(line5)
	DrawLine(line6)
	DrawLine(line7)
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
