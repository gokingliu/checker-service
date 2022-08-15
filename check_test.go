package main

import (
	"context"
	"reflect"
	"testing"

	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"github.com/golang/mock/gomock"

	pb "git.woa.com/crotaliu/pb-hub"
)

//go:generate go mod tidy

//go:generate mockgen -destination=stub/git.woa.com/crotaliu/pb-hub/checker_mock.go -package=pb_hub -self_package=git.woa.com/crotaliu/pb-hub --source=stub/git.woa.com/crotaliu/pb-hub/checker.trpc.go

func Test_checkImpl_Health(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	checkService := pb.NewMockCheckService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := checkService.EXPECT().Health(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HealthRequest, rsp *pb.HealthReply) error {
		s := &checkImpl{}
		return s.Health(ctx, req, rsp)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.HealthRequest
		rsp *pb.HealthReply
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp := &pb.HealthReply{}
			if err := checkService.Health(tt.args.ctx, tt.args.req, rsp); (err != nil) != tt.wantErr {
				t.Errorf("checkImpl.Health() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("checkImpl.Health() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
