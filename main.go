package main

import (
	"flag"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	w := flag.Duration("w", 25*time.Minute, "work duration")
	b := flag.Duration("b", 5*time.Minute, "break duration")
	flag.Parse()
	for {
		beeep.Alert("pomodoro", "Begin Pomodoro session", "")
		time.Sleep(*w)
		beeep.Alert("pomodoro", "Time for a break", "")
		time.Sleep(*b)
	}
}
