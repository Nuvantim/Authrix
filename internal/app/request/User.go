package request

import pg "github.com/jackc/pgx/v5/pgtype"

type UpdateAccount struct {
	Name     string  `validate:"required,min=2" json:"name"`
	Password string  `validate:"omitempty,min=1" json:"password"`
	Age      pg.Int4 `validate:"omitempty,gt=0" json:"age"`
	Phone    pg.Int4 `validate:"omitempty,gt=0" json:"phone"`
	District pg.Text `validate:"omitempty" json:"district"`
	City     pg.Text `validate:"omitempty" json:"city"`
	Country  pg.Text `validate:"omitempty" json:"country"`
}

type UpdateClient struct {
	Name     string  `validate:"required,min=2" json:"name"`
	Email    string  `validate:"omitempty,email" json:"email"`
	Password string  `validate:"omitempty,min=8" json:"password"`
	Role     []int32 `validate:"omitempty,dive,gt=0" json:"role_id"`
}
