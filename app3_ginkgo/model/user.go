package model

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
