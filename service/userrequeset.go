package service

type UserInfoGet struct {
	Id string `json:"id" form:"id" binding:"required" `
}
type UserListGet struct {
	Id        string `json:"id" form:"id" `
	Name      string `json:"name" form:"name"`
	PageNo    int    `json:"pageno" form:"pageno" binding:"required"`
	PageCount int    `json:"pagecount" form:"pagecount" binding:"required"`
}
