package models

type User struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}
