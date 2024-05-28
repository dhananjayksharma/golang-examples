package utils

import (
	"fmt"
	"time"
)

func GetCurrentTime() string {
	// Using time.Now() function.
	dt := time.Now()
	timeStr := fmt.Sprintf("Current date and time is: %s", dt.String())

	return timeStr
}
