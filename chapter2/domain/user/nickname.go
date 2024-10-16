package user

type Nickname struct {
	Value string `binding:"min=2,max=10"`
}
