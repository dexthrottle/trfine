package dto

type PostCreateDTO struct {
	Title      string `json:"title" form:"title" binding:"required"`
	Content    string `json:"content" form:"content" binding:"required"`
	UserID     uint64 `json:"-"`
	UserTgId   int    `json:"user_tg_id"`
	CategoryID uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
}
