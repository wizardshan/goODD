package request

type MobileFieldV2 struct {
	Mobile string `binding:"required,number,mobile"`
}

type CaptchaFieldV2 struct {
	Captcha string `binding:"required,number,len=4"`
}

type MobileField struct {
	Mobile string `binding:"number,mobile"`
}

type CaptchaField struct {
	Captcha string `binding:"number,len=4"`
}

type IDField struct {
	ID int64 `binding:"min=1"`
}
