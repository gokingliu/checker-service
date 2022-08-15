package main

import (
	"context"

	pb "git.woa.com/crotaliu/pb-hub"
)

type checkImpl struct{}

// Health 检查存活
func (s *checkImpl) Health(ctx context.Context, req *pb.HealthRequest, rsp *pb.HealthReply) error {
	// implement business logic here ...
	// ...
	return nil
}
