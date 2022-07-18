package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	t := time.Now()
	dayOffset := diff(int(t.Weekday()), 2)
	tz, _ := time.LoadLocation("America/New_York")
	showDate := time.Date(
		t.Year(),
		t.Month(),
		t.Day()+dayOffset,
		14,
		0,
		0,
		0,
		tz,
	)
	hours := time.Until(showDate).Hours()
	days := math.Floor(hours / 24)
	minutes := math.Ceil(math.Mod((hours-days), 1) * 60)
	if days > 1 {
		if days == 6 && hours == 23 {
			fmt.Printf("Live now! => <https://youtube.com/changelog>\n")
			os.Exit(0)
		}
		fmt.Printf("Next show in: %.0f days, %.0f hours, %.0f minutes\n", days, hours, minutes)
		os.Exit(0)
	}
	if hours > 1 {
		fmt.Printf("Next show in: %.0f hours, %.0f minutes\n", hours, minutes)
		os.Exit(0)
	}
	fmt.Printf("Grab a brew. Next show in %.0f minutes\n", minutes)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
