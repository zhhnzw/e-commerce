package utils

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Resp struct {
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
	Code    string      `json:"code"`
}

func ToJson(param interface{}) string {
	jsonValue, err := json.Marshal(param)
	zap.L().Error("json解析异常", zap.Error(err), zap.Any("param", param))
	return string(jsonValue)
}

// 常用的异常处理
func CheckErr(err error, msg string) {
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
	}
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func CheckRPCError(err error) error {
	if err != nil {
		s := status.Convert(err)
		err := errors.New(fmt.Sprintf("%v", s.Message()))
		switch s.Code() {
		case codes.InvalidArgument:
			log.Printf("RPC InvalidArgument: %v", s.Message())
			return err
		default:
			log.Printf("RPC Unexpected type: %v", s.Message())
			return err
		}
	}
	return nil
}
