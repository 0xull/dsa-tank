package leetcode

import (
	"strings"
	"testing"
)

func zigzagConversion(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	index := 0
	goingDown := true

	for _, r := range s {
		rows[index].WriteRune(r)

		switch index {
		case numRows - 1:
			goingDown = false
		case 0:
			goingDown = true
		}

		if goingDown {
			index++
		} else {
			index--
		}
	}

	var zigzagString strings.Builder
	for _, substr := range rows {
		zigzagString.WriteString(substr.String())
	}

	return zigzagString.String()
}

func TestZigzagConversion(t *testing.T) {

}
