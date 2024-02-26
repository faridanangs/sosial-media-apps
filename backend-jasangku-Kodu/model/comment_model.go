package model

import "time"

type Comment struct {
	User      User      `gorm:"foreignKey:id_user;references:id"`
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Comment   string    `gorm:"column:comment;not null"`
	IdUser    string    `gorm:"column:id_user;not null"`
	IdPost    string    `gorm:"column:id_post;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type CommentCreate struct {
	Comment string `json:"Comment" validate:"required"`
	IdUser  string `json:"IdUser" validate:"required"`
	IdPost  string `json:"IdPost" validate:"required"`
}

type CommentUpdate struct {
	Comment string `json:"Comment" validate:"required"`
}

type UserComment struct {
	ID       string `json:"id"`
	Image    string `json:"image"`
	Username string `json:"username"`
}

type CommentResponse struct {
	ID        int         `json:"id"`
	Comment   string      `json:"comment"`
	PostId    string      `json:"post_id"`
	CreatedAt time.Time   `json:"created_at"`
	User      UserComment `json:"user"`
}
