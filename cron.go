package main

import (
	"log"
	"time"

	"github.com/huanggengzhong/go-gin-example/models"
	"github.com/robfig/cron"
)

func main() {
	log.Println("starting....")
	c := cron.New()
	//cron表达式 "* * * * *" 表示时间表，格式为：分钟 小时 日期 月份 星期几。
	c.AddFunc("* * * * * *", func() {
		log.Println("开始CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("开始CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	// 	（2）for + select

	// 阻塞 select 等待 channel

	// （3）t1.Reset

	// 会重置定时器，让它重新开始计时
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}

}
