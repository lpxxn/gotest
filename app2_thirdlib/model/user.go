package model

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *UserInfo) GetName() string {
	return u.Name
}

func (u *UserInfo) SetName(name string) {
	u.Name = name
}

type UserRole struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
