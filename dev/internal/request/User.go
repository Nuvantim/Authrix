package request

type UpdateProfile struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      string `json:"age"`
	Phone    string `json:"phone"`
	District string `json:"district"`
	City     string `json:"city"`
	Country  string `json:"country"`
}
