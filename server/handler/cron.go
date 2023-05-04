package handler

import (
	"fmt"
	"subway/server/services/cart"
	"subway/server/services/dailyreward"
	"time"

	cron "github.com/robfig/cron/v3"
)

func StartCron() {
	fmt.Println("Cron is running")
	c := cron.New()
	c.AddFunc("*/2 * * * *", func() {
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
	c.AddFunc("*/2 * * * *", func() {
		fmt.Println("time is ,", time.Now())
		cart.GenerateCart()
	})

	c.Start()
}
