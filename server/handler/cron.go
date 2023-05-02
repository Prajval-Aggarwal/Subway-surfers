package handler

import (
	"fmt"
	"subway/server/services/dailyreward"
	"time"

	cron "github.com/robfig/cron/v3"
)

func StartCron() {
	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		fmt.Println("Time is:", time.Now())
		err := dailyreward.UpdateStreak()
		if err != nil {
			return
		}
		err = dailyreward.GenerateReward()
		if err != nil {
			return
		}
	})

	c.Start()
}
