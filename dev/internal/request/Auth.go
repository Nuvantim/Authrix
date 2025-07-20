package request

type OtpToken struct {
	Email string `json:"email"`
}

type Register struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassword struct {
	Code           string `json:"code"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
