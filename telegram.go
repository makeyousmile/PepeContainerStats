package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func StartBot() {
	pref := tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		text := "Pepe testnet container stats"
		return c.Send(text)
	})
	b.Handle("/stats", func(c tele.Context) error {
		stats := GetPepeContainerStats()
		cpuusage := fmt.Sprintf("%g", stats.CpuStats.Usage.Total)[:7]
		log.Print(cpuusage)
		text := "State: " + stats.State + "\n" + "Status: " + stats.Status + "\n" + "CPU usage: " + cpuusage
		return c.Send(text)
	})
	b.Handle("/logs", func(c tele.Context) error {
		stats := GetPepeContainerStats()
		text := GetContainerLogs(stats.Id)
		return c.Send(text)
	})
	b.Start()
}
