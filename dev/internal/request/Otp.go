package request

import "time"

type OtpToken struct {
	ID        int32     `json:"id"`
	Code      string    `json:"code"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
