package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"bytes"
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
	if widthInt == 0 {
		widthInt = 100
	}
	return widthInt
}


func RepeatChar(char string, num int) string{
	var buffer bytes.Buffer
    for i:=0;i<num;i++{
		buffer.WriteString(char)
	}
	return buffer.String()
}