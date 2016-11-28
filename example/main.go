package main

import (
	"github.com/goroom/aliyun_sms"
	"github.com/goroom/logger"
)

func main() {
	aliyun_sms, err := aliyun_sms.NewAliyunSms("签名名称", "模版CODE(SMS_*)", "Access Key ID", "Access Key Secret")
	if err != nil {
		logger.Error(err)
		return
	}
	err = aliyun_sms.Send("133********", `{"verifycode":"1234","minute":"30"}`)
	if err != nil {
		logger.Error(err)
		return
	}
}
