package dto

type CreateUserDTO struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Username  string `json:"username" binding:"required"`
	UserTGId  int    `json:"user_tg_id" binding:"required"`
}
