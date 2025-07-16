package request

import "time"

type OtpToken struct {
	ID        int32     `json:"id"`
	Code      string    `json:"code"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Register struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
