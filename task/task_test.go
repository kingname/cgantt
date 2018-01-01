package task

import (
	"testing"

	"github.com/kingname/cgantt/utils"
)

func Test_Drawline(t *testing.T) {
	width := utils.GetCurrentTerminalWidth()
	line := Line{"-", width / 2, "", 0, "=", width / 2}
	fullLine := DrawLine(line)
	answer := ""
	for i := 0; i < width/2; i++ {
		answer += "-"
	}
	for i := 0; i < width/2; i++ {
		answer += "="
	}
	if answer == fullLine {
		t.Log("Pass")
	} else {
		t.Error("Draw Line Error.")
	}
}
