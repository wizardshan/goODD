package vo

type Captcha struct {
	Value string `binding:"number,len=4"`
}
