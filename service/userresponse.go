package service

type UserInfoGetRes struct {
	Id       string      `json:"id" sql:"userid"`
	Name     string      `json:"name" sql:"username"`
	Level    int         `json:"level" sql:"secretlevel"`
	Email    string      `json:"email" sql:"email"`
	Phonenum string      `json:"phonenum" sql:"lxdh"`
	userinfo UserInfoGet `json:"userinfo"`
	roleid   []string    `json:"roleid" `
}

type UserListGetRes struct {
	Id       string `json:"id" sql:"userid"`
	Name     string `json:"name" sql:"username"`
	Level    int    `json:"level" sql:"secretlevel"`
	Email    string `json:"email" sql:"email"`
	Phonenum string `json:"phonenum" sql:"lxdh"`
}
