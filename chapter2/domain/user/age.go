package user

type Age struct {
	Value int `binding:"min=1,max=120"`
}
