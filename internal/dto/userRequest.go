package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerifyUser struct {
	Code int `json:"code"`
}
