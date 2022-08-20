package services

import (
	"checker/configs"
	"checker/logic"
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
)

type CheckImpl struct{}

// Health 检查存活
func (s *CheckImpl) Health(ctx context.Context, req *pb.HealthRequest, rsp *pb.HealthReply) error {
	// 检查进程和文件
	if req.Type == 1 {
		// 检查进程
		processResult, processFailStr, processErr := logic.GetHealthLogic("process")
		// 读取或解析 yaml 错误
		if processErr != nil {
			rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
			rsp.Result = false
			return nil
		}
		// 返回检查结果
		if processFailStr != "" {
			rsp.Msg = processFailStr
		} else {
			rsp.Msg = configs.ResOk.Msg
		}
		rsp.Code = configs.ResOk.Code
		rsp.Result = processResult

		return nil
	} else if req.Type == 2 {
		// 检查文件
		filesResult, filesFailStr, filesErr := logic.GetHealthLogic("files")
		// 读取或解析 yaml 错误
		if filesErr != nil {
			rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
			rsp.Result = false
			return nil
		}
		// 返回检查结果
		if filesFailStr != "" {
			rsp.Msg = filesFailStr
		} else {
			rsp.Msg = configs.ResOk.Msg
		}
		rsp.Code = configs.ResOk.Code
		rsp.Result = filesResult

		return nil
	} else {
		// 检查进程
		processResult, processFailStr, processErr := logic.GetHealthLogic("process")
		// 检查文件
		filesResult, filesFailStr, filesErr := logic.GetHealthLogic("files")
		// 读取或解析 yaml 错误
		if processErr != nil || filesErr != nil {
			rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
			rsp.Result = false
			return nil
		}
		// 返回检查结果
		if processFailStr != "" || filesFailStr != "" {
			rsp.Msg = processFailStr + filesFailStr
		} else {
			rsp.Msg = configs.ResOk.Msg
		}
		rsp.Code = configs.ResOk.Code
		rsp.Result = processResult && filesResult

		return nil
	}
}
