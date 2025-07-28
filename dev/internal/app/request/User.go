package request

import "github.com/jackc/pgx/v5/pgtype"

type UpdateAccount struct {
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Age      pgtype.Int4 `json:"age"`
	Phone    pgtype.Int4 `json:"phone"`
	District pgtype.Text `json:"district"`
	City     pgtype.Text `json:"city"`
	Country  pgtype.Text `json:"country"`
}

type UpdateClient struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"Password"`
}
