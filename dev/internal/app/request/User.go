package request

import pg "github.com/jackc/pgx/v5/pgtype"

type UpdateAccount struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Age      pg.Int4 `json:"age"`
	Phone    pg.Int4 `json:"phone"`
	District pg.Text `json:"district"`
	City     pg.Text `json:"city"`
	Country  pg.Text `json:"country"`
}

type UpdateClient struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Role     []int32 `json:"role_id"`
}
