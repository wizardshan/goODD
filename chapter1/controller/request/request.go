package request

type MobileFieldV2 struct {
	Mobile string `binding:"required,number,mobile"`
}

type MobileField struct {
	Mobile string `binding:"number,mobile"`
}

type IDField struct {
	ID int64 `binding:"min=1"`
}
