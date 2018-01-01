package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func raiseErr(err error) {
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// getCurrentTerminalWidth do what its name represents it will do
func GetCurrentTerminalWidth() int {
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
