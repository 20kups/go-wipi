package utils

import (
	"fmt"
	"time"
)

var nowFmtStr = "Monday, January 2, 2006 at 3:04:05 PM"

func PrintlnWithTimestamp(msg string) {
	fmt.Println(SprintWithTimestamp(msg))
}
func PrintSameLineWithTimestamp(msg string) {
	fmt.Print("\r\033[2K")
	fmt.Printf(SprintWithTimestamp(msg))
}

func SprintWithTimestamp(msg string) string {
	// Get current date and time formatted alongside the message
	return fmt.Sprintf("[%v] %s", time.Now().Format(nowFmtStr), msg)
}