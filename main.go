package main

import (
	"github.com/docker/docker/client"
	"log"
)

type Cfg struct {
	cli      *client.Client
	BotToken string
}

var cfg = &Cfg{}

func init() {
	cfg.cli = NewDockerClient()
	cfg.BotToken = "6336512021:AAEvm9RFs1yLiiJmicxBLZJOXLixcLRqusQ"
}
func main() {

	stats := GetPepeContainerStats()
	log.Print(stats.Status)
	StartBot()

}
