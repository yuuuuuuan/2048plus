package game2048

const (
	msgTitle     = "2048"
	msgCopyright = "made by yuuuuuuan"
	msgStart     = "游戏开始"
	msgExit      = "退出游戏"
)

type Profile struct {
	button1   *Button
	button2   *Button
	title     string
	copyright string
}

func NewProfile() *Profile {
	button1 := NewButton(msgStart, ScreenWidth/2, ScreenHeight/2)
	button2 := NewButton(msgExit, ScreenWidth/2, ScreenHeight/2-buttonHeight*2)

	p := &Profile{
		button1:   button1,
		button2:   button2,
		title:     msgTitle,
		copyright: msgCopyright,
	}
	return p
}
