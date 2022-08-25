package services

import (
	"checker/configs"
	"checker/logic"
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
	"strings"
)

type CheckImpl struct{}

// Check 检查脚本存活
func (s *CheckImpl) Check(ctx context.Context, req *pb.HealthRequest, rsp *pb.HealthReply) error {
	// 检查脚本
	scriptsResult, scriptsFailStr, scriptsErr := logic.CheckLogic("scripts")
	// 读取或解析 yaml 错误
	if scriptsErr != nil {
		rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
		rsp.Result = false
		return nil
	}
	// 返回检查结果
	if scriptsFailStr != "" {
		rsp.Code, rsp.Msg = configs.ResFail.Code, scriptsFailStr
	} else {
		rsp.Code, rsp.Msg = configs.ResOk.Code, configs.ResOk.Msg
	}
	rsp.Result = scriptsResult

	return nil
}

// Health 检查进程和文件存活
func (s *CheckImpl) Health(ctx context.Context, req *pb.HealthRequest, rsp *pb.HealthReply) error {
	// 检查进程和文件
	if req.Type == 1 {
		// 检查进程
		processResult, processFailStr, processErr := logic.HealthLogic("process")
		// 读取或解析 yaml 错误
		if processErr != nil {
			rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
			rsp.Result = false
			return nil
		}
		// 返回检查结果
		if processFailStr != "" {
			rsp.Code, rsp.Msg = configs.ResFail.Code, processFailStr
		} else {
			rsp.Code, rsp.Msg = configs.ResOk.Code, configs.ResOk.Msg
		}
		rsp.Result = processResult

		return nil
	} else if req.Type == 2 {
		// 检查文件
		filesResult, filesFailStr, filesErr := logic.HealthLogic("files")
		// 读取或解析 yaml 错误
		if filesErr != nil {
			rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
			rsp.Result = false
			return nil
		}
		// 返回检查结果
		if filesFailStr != "" {
			rsp.Code, rsp.Msg = configs.ResFail.Code, filesFailStr
		} else {
			rsp.Code, rsp.Msg = configs.ResOk.Code, configs.ResOk.Msg
		}
		rsp.Result = filesResult

		return nil
	} else {
		// 检查进程
		processResult, processFailStr, processErr := logic.HealthLogic("process")
		// 检查文件
		filesResult, filesFailStr, filesErr := logic.HealthLogic("files")
		// 读取或解析 yaml 错误
		if processErr != nil || filesErr != nil {
			rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
			rsp.Result = false
			return nil
		}
		// 返回检查结果
		if processFailStr != "" || filesFailStr != "" {
			rsp.Code = configs.ResFail.Code
			if processFailStr != "" && filesFailStr != "" {
				rsp.Msg = processFailStr + "\n" + filesFailStr
			} else {
				rsp.Msg = processFailStr + filesFailStr
			}
		} else {
			rsp.Code, rsp.Msg = configs.ResOk.Code, configs.ResOk.Msg
		}
		rsp.Result = processResult && filesResult

		return nil
	}
}

func (s *CheckImpl) GetHealth(ctx context.Context, req *pb.HealthRequest, rsp *pb.HealthReply) error {
	msgErr, err := logic.GetHealthLogic(req.Type)
	// 解析出错
	if err != nil {
		rsp.Code, rsp.Msg = configs.InnerUnmarshalYamlError.Code, configs.InnerUnmarshalYamlError.Msg
		rsp.Result = false
		return nil
	}
	// 机器出错信息
	if len(msgErr) > 0 {
		rsp.Code, rsp.Msg = configs.ResFail.Code, strings.Join(msgErr, "\n")
		rsp.Result = false
		return nil
	}
	rsp.Code, rsp.Msg = configs.ResOk.Code, configs.ResOk.Msg
	rsp.Result = true

	return nil
}
