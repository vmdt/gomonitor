package client

import (
	"context"
	"io"
	"log"
	"producer/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct{}

func (c *Client) FetchStats(s proto.StatsServiceClient) (map[string]int, error) {
	var stats = map[string]int{}
	strem, err := s.FetchStats(context.Background(), &emptypb.Empty{})
	if err == io.EOF {
		log.Println("EOF")
	}
	if err != nil {
		return nil, err
	}

	for {
		resp, recvErr := strem.Recv()
		if recvErr == io.EOF {
			log.Println("server finish streaming")
			return stats, nil
		}
		if resp.Bandwidth != nil {
			stats[resp.MemTag] = int(*resp.Bandwidth)
		} else {
			stats[resp.MemTag] = 0
		}
	}
}
