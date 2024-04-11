package main

import (
	"fmt"
	"log"
	"net/http"
	grpc_client "producer/client"
	"producer/config"
	"producer/handlers"
	"producer/proto"

	"github.com/robfig/cron"
	"google.golang.org/grpc"
)

func main() {
	cf := config.NewConfig()

	conn, err := grpc.Dial("localhost:50069", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error while dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewStatsServiceClient(conn)
	runCronStatsJob(client)

	http.ListenAndServe(fmt.Sprintf(":%s", cf.PRODUCER_PORT), nil)
}

func runCronStatsJob(client proto.StatsServiceClient) {
	gclient := grpc_client.Client{}

	c := cron.New()
	c.AddFunc("@every 0h0m3s", func() {
		stats, err := gclient.FetchStats(client)
		err = handlers.PushDataToKafka(stats)

		if err != nil {
			log.Printf("error while fetch stats: %v", err)
			c.Stop()
			return
		}
	})

	c.Start()
}
