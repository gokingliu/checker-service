package services

import (
	"checker/logic"
	"context"
	"git.code.oa.com/trpc-go/trpc-go/log"
	pb "git.woa.com/crotaliu/pb-hub"
	"strings"
)

type CheckImpl struct{}

// Health 检查存活
func (s *CheckImpl) Health(ctx context.Context, req *pb.HealthRequest, rsp *pb.HealthReply) error {
	if req == nil || req.Type == 0 {
		processResult, processFailList, err := logic.GetHealthLogic("process")
		if err != nil {
			return err
		}
		stringY := strings.Join(processFailList, ",")
		log.Error(stringY)

		filesResult, filesFailList, err := logic.GetHealthLogic("files")
		if err != nil {
			return err
		}
		stringX := strings.Join(filesFailList, ",")
		log.Error(stringX)

		rsp.IsAlive = processResult && filesResult
	}

	return nil
}
