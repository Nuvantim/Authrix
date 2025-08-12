package request

type OtpToken struct {
	Email string `validate:"required,email" json:"email"`
}

type Register struct {
	Code     string `validate:"required" json:"code"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}

type ResetPassword struct {
	Code           string `validate:"required" json:"code"`
	Password       string `validate:"required,min=8" json:"password"`
	RetypePassword string `validate:"required,eqfield=Password" json:"retype_password"`
}

type Login struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=1" json:"password"`
}
