package dto

type CreateCategoryDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	UserID   uint64 `json:"-"`
	UserTgId int    `json:"user_tg_id"`
}
