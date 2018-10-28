package main

import (
	"flag"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	w := flag.Duration("w", 25*time.Minute, "work duration")
	b := flag.Duration("b", 5*time.Minute, "break duration")
	flag.Parse()
	startMessage, breakMessage := "Begin Pomodoro session", "Time for a break"
	if m := os.Getenv("START_MESSAGE"); m != "" {
		startMessage = m
	}
	if m := os.Getenv("BREAK_MESSAGE"); m != "" {
		breakMessage = m
	}
	for {
		beeep.Alert("pomodoro", startMessage, "")
		time.Sleep(*w)
		beeep.Alert("pomodoro", breakMessage, "")
		time.Sleep(*b)
	}
}
