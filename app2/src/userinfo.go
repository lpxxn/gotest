package src

type UserInfo struct {
	Name string
	Age  int
}


func (u *UserInfo) GetName() string {
	return u.Name
}

func (u *UserInfo) SetName(name string) {
	u.Name = name
}

func NewUserInfo() *UserInfo {
	return &UserInfo{Name:RandomStr(3), Age:RandomInt(5, 30)}
}
