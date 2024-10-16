package user

type AgeField struct {
	Age int `binding:"min=1,max=120"`
}

type LevelField struct {
	Level int `binding:"oneof=0 10 20 30"`
}

type NicknameField struct {
	Nickname string `binding:"min=2,max=10"`
}
