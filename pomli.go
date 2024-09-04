package main

import (
	"flag"
	"fmt"
	"pomli/internal/timer"
	"time"
)

func main() {
	minutes := flag.Bool("m", false, "set timer in minutes, default = seconds")
	duration := flag.Int("t", 10, "duration of timer")
	help := flag.Bool("h", false, "display help message")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *minutes {
		*duration = *duration * int(time.Minute)
	} else {
		*duration = *duration * int(time.Second)
	}

	newTimer, err := timer.CreateTimer(*duration)

	if err != nil {
		fmt.Println("Error:", err)
	}

	timer.ExecuteTimer(*newTimer)
}
