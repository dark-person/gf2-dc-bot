package main

import (
	"fmt"
	"time"
)

func main() {
	// Start up log
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Starting up...")

	// Init task
	initDailyTask()

	// Keep the main program running
	select {}
}
