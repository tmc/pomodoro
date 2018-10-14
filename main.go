package main

import (
	"flag"
	"time"

	"github.com/gen2brain/beeep"
)

var (
	flagWorkDuration  = flag.Duration("w", 25*time.Minute, "work duration")
	flagBreakDuration = flag.Duration("b", 5*time.Minute, "break duration")
)

func main() {
	flag.Parse()
	for {
		beeep.Alert("pomodoro", "Begin Pomodoro session", "")
		time.Sleep(*flagWorkDuration)
		beeep.Alert("pomodoro", "Time for a break", "")
		time.Sleep(*flagBreakDuration)
	}
}
