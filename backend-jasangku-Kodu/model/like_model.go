package model

import "time"

type Like struct {
	IdUser    string    `gorm:"column:id_user;not null"`
	IdPost    string    `gorm:"column:id_post;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

type LikeCreate struct {
	IdUser string `json:"IdUser" validate:"required"`
	IdPost string `json:"IdPost" validate:"required"`
}

type LikeResponse struct {
	ID        int       `json:"id"`
	IdUser    string    `json:"id_user"`
	IdPost    string    `json:"id_post"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Like) TableName() string {
	return "mtm_posts_users"
}
