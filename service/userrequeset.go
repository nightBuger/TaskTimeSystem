package service

type UserInfoGet struct {
	Id string `json:"id" form:"id" binding:"required"`
}
