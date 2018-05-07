package main

import (
	"fmt"

	"github.com/goroom/aliyun_sms"
)

func main() {
	aliyun_sms, err := aliyun_sms.NewAliyunSms("****", "SMS_****", "LTA*****vBuaN", "bSduBzta************gfxkXM")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = aliyun_sms.Send("13319****73", `{"VerifyCode":"1234","Minutes":"30"}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
