package main

import (
	"github.com/goroom/aliyun_sms"
	"github.com/goroom/logger"
)

func main() {
	aliyun_sms, err := aliyun_sms.NewAliyunSms("****", "SMS_****", "LTA*****vBuaN", "bSduBzta************gfxkXM")
	if err != nil {
		logger.Error(err)
		return
	}
	err = aliyun_sms.Send("13319****73", `{"VerifyCode":"1234","Minutes":"30"}`)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Success")
}
