package gowx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

//https://mp.weixin.qq.com/debug/wxadoc/dev/api/api-login.html#wxloginobject
//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code

const (
	WxMina_URL_JsCode2Session = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type WxMina struct {//微信小程序
	Appid string
	Secret string
}

type WxMinaErr struct {
	ErrCode int `json:"errcode,omitempty"`
	ErrMsg string `json:"errmsg,omitempty"`
}

type WxMinaRspJsCode2Session struct {
	WxMinaErr
	OpenId string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Unionid string `json:"unionid,omitempty"`
}

var gWxMina *WxMina

func InitWxMina(appid string, secret string) {
	gWxMina = &WxMina{
		Appid: appid,
		Secret: secret,
	}
}

func GetWxMina() *WxMina {
	return gWxMina
}

/*
//正常返回的JSON数据包
{
      "openid": "OPENID",
      "session_key": "SESSIONKEY",
      "unionid": "UNIONID"
}
//错误时返回JSON数据包(示例为Code无效)
{
    "errcode": 40029,
    "errmsg": "invalid code"
}
*/

//first function 类型语言
//c#,java  frist class

//开发者服务器使用登录凭证 code 获取 session_key 和 openid
//session_key 是对用户数据进行加密签名的密钥。为了自身应用安全，session_key 不应该在网络上传输。
func (this *WxMina) JsCode2Session(jsCode string) (*WxMinaRspJsCode2Session, error){
	
	resp, err := http.Get(fmt.Sprintf(WxMina_URL_JsCode2Session, this.Appid, this.Secret,jsCode))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result WxMinaRspJsCode2Session
	
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, err
    } else {
        return &result,nil
    }
}