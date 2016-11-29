package aliyun_sms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

type Request struct {
	//public
	Format           string //否	返回值的类型，支持JSON与XML。默认为XML
	Version          string //是	API版本号，为日期形式：YYYY-MM-DD，本版本对应为2016-09-27
	AccessKeyId      string //是	阿里云颁发给用户的访问服务所用的密钥ID
	Signature        string //是	签名结果串，关于签名的计算方法，请参见 签名机制。
	SignatureMethod  string //是	签名方式，目前支持HMAC-SHA1
	Timestamp        string //是	请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为YYYY-MM-DDThh:mm:ssZ 例如，2015-11-23T04:00:00Z（为北京时间2015年11月23日12点0分0秒）
	SignatureVersion string //是	签名算法版本，目前版本是1.0
	SignatureNonce   string //是	唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值

	//sms
	Action       string //必须	操作接口名，系统规定参数，取值：SingleSendSms
	SignName     string //必须	管理控制台中配置的短信签名（状态必须是验证通过）
	TemplateCode string //必须	管理控制台中配置的审核通过的短信模板的模板CODE（状态必须是验证通过）
	RecNum       string //必须	目标手机号，多个手机号可以逗号分隔
	ParamString  string //必选	短信模板中的变量；数字需要转换为字符串；个人用户每个变量长度必须小于15个字符。 例如:短信模板为：“接受短信验证码${no}”,此参数传递{“no”:”123456”}，用户将接收到[短信签名]接受短信验证码123456
}

func (this *Request) signString(source string, access_secret string) string {
	//fmt.Println("source:", source)
	//fmt.Println("access_secret:", access_secret)
	h := hmac.New(sha1.New, []byte(access_secret))
	h.Write([]byte(source))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (this *Request) ComputeSignature(params map[string](string), method string, access_secret string) string {
	var list []string
	for k, v := range params {
		list = append(list, this.PercentEncode(k)+"="+url.QueryEscape(v))
	}
	sort.Strings(list)

	canonicalized_query_string := strings.Join(list, "&")
	string_to_sign := method + "&%2F&" + this.PercentEncode(canonicalized_query_string)
	return this.signString(string_to_sign, access_secret+"&")
}

func (this *Request) ComposeUrl(method string, access_secret string) string {
	var params map[string](string)
	params = make(map[string](string))
	params["Format"] = this.Format
	params["Version"] = this.Version
	params["AccessKeyId"] = this.AccessKeyId
	params["SignatureMethod"] = this.SignatureMethod
	params["Timestamp"] = this.Timestamp
	params["SignatureVersion"] = this.SignatureVersion
	params["SignatureNonce"] = this.SignatureNonce
	params["Action"] = this.Action
	params["SignName"] = this.SignName
	params["TemplateCode"] = this.TemplateCode
	params["RecNum"] = this.RecNum
	params["ParamString"] = this.ParamString

	params["Signature"] = this.ComputeSignature(params, method, access_secret)
	//fmt.Println("Signature:", params["Signature"])

	var _url string
	if method == "POST" {
		_url = "http://sms.aliyuncs.com/"
	} else {
		_url = "http://sms.aliyuncs.com/?"
		for k, v := range params {
			//fmt.Println(k, v, url.QueryEscape(v))
			_url += k + "=" + url.QueryEscape(v) + "&"
		}
		_url = _url[:len(_url)-1]
	}

	return _url
}

func (this *Request) PercentEncode(str string) string {
	str = url.QueryEscape(str)

	return str
}
