package grpc

import (
	"stats-service/proto"
	"stats-service/services"

	"github.com/golang/protobuf/ptypes/empty"
)

type GRPCStatsService struct {
	svc services.StatsService
	proto.UnimplementedStatsServiceServer
}

func NewGRPCStatsService(svc services.StatsService) *GRPCStatsService {
	return &GRPCStatsService{
		svc: svc,
	}
}

func (s *GRPCStatsService) FetchStats(_ *empty.Empty, stream proto.StatsService_FetchStatsServer) error {
	stats, err := s.svc.ReadMemory()
	if err != nil {
		return err
	}

	for key, value := range stats {
		err := stream.Send(&proto.StatsResponse{
			MemTag:    key,
			Bandwidth: int32(value),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
