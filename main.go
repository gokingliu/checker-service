package main

import (
	"checker/services"
	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	_ "git.code.oa.com/trpc-go/trpc-filter/recovery"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/log"
	_ "git.code.oa.com/trpc-go/trpc-log-atta"
	_ "git.code.oa.com/trpc-go/trpc-metrics-m007"
	_ "git.code.oa.com/trpc-go/trpc-metrics-runtime"
	_ "git.code.oa.com/trpc-go/trpc-naming-polaris"
	pb "git.woa.com/crotaliu/pb-hub"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterCheckService(s, &services.CheckImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
