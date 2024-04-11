package main

import (
	"fmt"
	"log"
	"net"
	grpc_server "stats-service/grpc"
	"stats-service/proto"
	"stats-service/services"

	"google.golang.org/grpc"
)

func main() {
	var (
		config = NewConfig()
		svc    = &services.Stats{}
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.STATS_PORT))
	if err != nil {
		log.Fatalf("error while creating grpc server: %v", err)
	}

	s := grpc.NewServer()
	grpcStatsService := grpc_server.NewGRPCStatsService(svc)
	proto.RegisterStatsServiceServer(s, grpcStatsService)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
