package main

import (
	"context"

	pb "git.code.oa.com/trpcprotocol/test/helloworld"
)

type greeterImpl struct{}

func (s *greeterImpl) SayHello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) error {
	// implement business logic here ...
	// ...
	return nil
}
