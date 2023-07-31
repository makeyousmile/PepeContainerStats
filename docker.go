package main

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"log"
	"time"
)

type ContainerStats struct {
	Id       string `json:"id"`
	Read     string `json:"read"`
	Preread  string `json:"preread"`
	CpuStats cpu    `json:"cpu_stats"`
	State    string
	Status   string
}

type cpu struct {
	Usage cpuUsage `json:"cpu_usage"`
}

type cpuUsage struct {
	Total float64 `json:"total_usage"`
}

func NewDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	return cli
}

func GetContainerList() []types.Container {

	containers, err := cfg.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return containers
}

func GetPepeContainerStats() ContainerStats {

	containers := GetContainerList()

	for _, container := range containers {
		//learningathome/petals
		if container.Image == "test" {
			containerStats := GetContainerStats(container.ID)
			containerStats.State = container.State
			containerStats.Status = container.Status
			containerStats.Id = container.ID
			return containerStats
		}
	}

	return ContainerStats{}
}
func GetContainerLogs(id string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reader, err := cfg.cli.ContainerLogs(ctx, id, types.ContainerLogsOptions{ShowStdout: true})
	logs, err := io.ReadAll(reader)
	if err != nil {
		log.Print(err)
	}
	return string(logs)
}

func GetContainerStats(id string) ContainerStats {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stats, err := cfg.cli.ContainerStatsOneShot(ctx, id)
	if err != nil {
		log.Print(err)
	}
	decoder := json.NewDecoder(stats.Body)
	var containerStats ContainerStats
	err = decoder.Decode(&containerStats)
	return containerStats
}
