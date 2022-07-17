package service

type UserInfoGetRes struct {
	Id    string `json:"id" sql:"userid"`
	Name  string `json:"name" sql:"username"`
	Email string `json:"email" sql:"email"`
	Tel   string `json:"tel" sql:"tel"`
}

type UserListGetRes struct {
	Id    string `json:"id" sql:"userid"`
	Name  string `json:"name" sql:"username"`
	Email string `json:"email" sql:"email"`
	Tel   string `json:"tel" sql:"tel"`
}
