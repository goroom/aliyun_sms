package aliyun_sms

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

type Request struct {
	//public
	Format      string //否	返回值的类型，支持JSON与XML。默认为XML
	Version     string //是	API版本号，为日期形式：YYYY-MM-DD，本版本对应为2016-09-27
	AccessKeyId string //是	阿里云颁发给用户的访问服务所用的密钥ID
	//Signature        string //是	签名结果串，关于签名的计算方法，请参见 签名机制。
	SignatureMethod  string //是	签名方式，目前支持HMAC-SHA1
	Timestamp        string //是	请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为YYYY-MM-DDThh:mm:ssZ 例如，2015-11-23T04:00:00Z（为北京时间2015年11月23日12点0分0秒）
	SignatureVersion string //是	签名算法版本，目前版本是1.0
	SignatureNonce   string //是	唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值

	//sms
	Action       string //必须	操作接口名，系统规定参数，取值：SendSms
	SignName     string //必须	管理控制台中配置的短信签名（状态必须是验证通过）
	TemplateCode string //必须	管理控制台中配置的审核通过的短信模板的模板CODE（状态必须是验证通过）
	PhoneNumbers string //必须	目标手机号，多个手机号可以逗号分隔
	RegionId      string //必须 API支持的RegionID，如短信API的值为：cn-hangzhou
	TemplateParam string //必选	短信模板中的变量；数字需要转换为字符串；个人用户每个变量长度必须小于15个字符。 例如:短信模板为：“接受短信验证码${no}”,此参数传递{“no”:”123456”}，用户将接收到[短信签名]接受短信验证码123456
}

// signString 用指定的access_secret 对source进行签名
func (this *Request) signString(source string, access_secret string) string {
	h := hmac.New(sha1.New, []byte(access_secret))
	h.Write([]byte(source))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// ComputeSignature 生成签名
func (this *Request) ComputeSignature(sortQueryString string, access_secret string) string {
	var popBuf bytes.Buffer
	popBuf.WriteString("GET&")
	popBuf.WriteString(specialURLEncode("/"))
	popBuf.WriteString("&")
	popBuf.WriteString(specialURLEncode(sortQueryString))
	return this.signString(popBuf.String(), access_secret+"&")
}

func (this *Request) ComposeUrl(method string, access_secret string) string {
	preSingURL := url.Values{}
	preSingURL.Add("AccessKeyId", this.AccessKeyId)
	preSingURL.Add("Action", this.Action)
	preSingURL.Add("Format", this.Format)
	preSingURL.Add("PhoneNumbers", this.PhoneNumbers)
	preSingURL.Add("RegionId", this.RegionId)
	preSingURL.Add("SignName", this.SignName)
	preSingURL.Add("SignatureMethod", this.SignatureMethod)
	preSingURL.Add("SignatureNonce", this.SignatureNonce)
	preSingURL.Add("SignatureVersion", this.SignatureVersion)
	preSingURL.Add("TemplateCode", this.TemplateCode)
	preSingURL.Add("TemplateParam", this.TemplateParam)
	preSingURL.Add("Timestamp", this.Timestamp)
	preSingURL.Add("Version", this.Version)
	sortStr := sortQueryString(preSingURL)
	Signature := specialURLEncode(this.ComputeSignature(sortStr, access_secret))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + sortStr

	return _url
}

func sortQueryString(preSingURL url.Values) string {
	var buffer bytes.Buffer
	keys := make([]string, 0, len(preSingURL))
	for k := range preSingURL {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if buffer.Len() > 0 {
			buffer.WriteString("&")
		}
		buffer.WriteString(specialURLEncode(k))
		buffer.WriteString("=")
		buffer.WriteString(specialURLEncode(preSingURL.Get(k)))
	}
	return buffer.String()
}

func specialURLEncode(str string) string {
	str = url.QueryEscape(str)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}
