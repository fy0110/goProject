package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}
type LoginResMes struct {
	Code  int //返回状态码 500 表示用户未注册 200登录成功
	Error string
}
