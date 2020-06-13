package controller

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"net/http"
)

// 微信相关接口

type WxController struct {
	Ctx iris.Context
	Visits uint64
}

func (c *WxController) GetWxuserinfo()  {
	js_code := c.Ctx.URLParamDefault("js_code","")
	if js_code == "" {
		c.Ctx.JSON("{code:-1,msg:'缺少参数'}")
	}else{
		var url bytes.Buffer
		url.WriteString("https://api.weixin.qq.com/sns/jscode2session")
		url.WriteString("?appid=")
		url.WriteString("wx03ae5b075cd517ed")
		url.WriteString("&secret=")
		url.WriteString("343026fd75e01d9db819560c17cc75eb")
		url.WriteString("&js_code=")
		url.WriteString(js_code)
		url.WriteString("&grant_type=authorization_code")
		resp, err := http.Get(url.String())
		if err != nil {
			log.Println("请求异常：", err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		if err != nil {
			// handle error
		}
		c.Ctx.StatusCode(iris.StatusOK)
		c.Ctx.WriteString(string(body))
	}
}