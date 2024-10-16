package user

type Level struct {
	Value int `binding:"oneof=0 10 20 30"`
}
