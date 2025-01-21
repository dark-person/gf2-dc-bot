package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Use https://crontab.guru/ to verify

// Create cron task that work daily.
func initDailyTask() {
	c := cron.New()

	c.AddFunc("0 9 * * *", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Send testing message.")

	})

	c.Start()
}
