package aliyun_sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/goroom/rand"
)

type AliyunSms struct {
	AccessKeyID  string
	AccessSecret string

	SignName     string //签名名称
	TemplateCode string //模板code
}

func NewAliyunSms(sign_name string, template_code string, access_key_id string, access_secret string) (*AliyunSms, error) {
	var a AliyunSms
	a.SignName = sign_name
	a.TemplateCode = template_code
	a.AccessKeyID = access_key_id
	a.AccessSecret = access_secret

	return &a, nil
}

func (this *AliyunSms) Send(numbers string, params string) error {
	var request Request
	request.Format = "JSON"
	request.Version = "2016-09-27"
	request.AccessKeyId = this.AccessKeyID
	request.SignatureMethod = "HMAC-SHA1"
	request.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	request.SignatureVersion = "1.0"
	request.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)

	request.Action = "SingleSendSms"
	request.SignName = this.SignName
	request.TemplateCode = this.TemplateCode
	request.RecNum = numbers
	request.ParamString = params

	url := request.ComposeUrl("GET", this.AccessSecret)
	fmt.Println(url)
	var resp *http.Response
	var err error
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_m := make(map[string](string))
	err = json.Unmarshal(b, &_m)
	if err != nil {
		return err
	}
	message, ok := _m["Message"]
	if ok {
		return errors.New(message)
	}

	return nil
}
