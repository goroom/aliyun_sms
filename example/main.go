package main

import (
	"github.com/goroom/aliyun_sms"
	"github.com/goroom/logger"
)

func main() {
	aliyun_sms, err := aliyun_sms.NewAliyunSms("巨灵易停", "SMS_27880074", "LTAIR8r5E68vBuaN", "bSduBztakdjOuEq5HF0199tBgfxkXM")
	if err != nil {
		logger.Error(err)
		return
	}
	err = aliyun_sms.Send("13319257173", `{"VerifyCode":"1234","Minutes":"30"}`)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Success")
}
