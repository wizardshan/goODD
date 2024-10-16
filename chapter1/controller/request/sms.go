package request

type SmsCaptchaV1 struct {
	Mobile string `binding:"required,number,mobile"`
}

type SmsCaptchaV2 struct {
	MobileFieldV2
}

type SmsCaptcha struct {
	MobileField `binding:"required"`
}
