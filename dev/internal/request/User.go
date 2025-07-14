package request

import "time"

type UserAccount struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserProfile struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Age       string    `json:"age"`
	Phone     string    `json:"phone"`
	District  string    `json:"district"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
