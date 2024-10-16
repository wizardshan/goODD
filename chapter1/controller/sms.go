package controller

import (
	"github.com/gin-gonic/gin"
	"goODD/chapter1/controller/request"
	"goODD/chapter1/controller/response"
)

type Sms struct {
}

func NewSms() *Sms {
	ctr := new(Sms)
	return ctr
}

func (ctr *Sms) CaptchaV1(c *gin.Context) (response.Data, error) {
	req := new(request.SmsCaptchaV1)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *Sms) CaptchaV2(c *gin.Context) (response.Data, error) {
	req := new(request.SmsCaptchaV2)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
	req := new(request.SmsCaptcha)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}
