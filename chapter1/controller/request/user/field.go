package user

type NicknameField struct {
	Nickname string `binding:"min=2,max=10"`
}

type LevelField struct {
	Level int `binding:"oneof=0 10 20 30"`
}
