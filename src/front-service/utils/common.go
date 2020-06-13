package utils

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
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
	Logf(err, "json解析异常:%+v", param)
	return string(jsonValue)
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func Logf(err error, format string, args ...interface{}) { // 打异常日志
	log.Printf("%+v", errors.Wrapf(err, format, args...))
}

func CheckErr(err error, format string, args ...interface{}) { // 异常处理
	if err != nil {
		log.Printf("%+v", errors.Wrapf(err, format, args...))
	}
}

func Fatalf(err error, format string, args ...interface{}) { // 异常处理
	if err != nil {
		log.Fatalf("%+v", errors.Wrapf(err, format, args...))
	}
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
