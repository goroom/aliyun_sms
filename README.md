# aliyun-sms

### Specified signature is not matched with our calculation
##### 参数填错可能也会报这个错误，并不是BUG
##### aliyun_sms.Send("133********", `{"verifycode":"1234","minute":"30"}`)
##### 第二个参数是json格式字符串，逗号后不能有空格，比如`{"verifycode":"1234", "minute":"30"}`是错误的，也会报签名错误。
