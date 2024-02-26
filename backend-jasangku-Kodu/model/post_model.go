package model

import "time"

type Post struct {
	User      User      `gorm:"foreignKey:id_user;references:id"`
	Comments  []Comment `gorm:"foreignKey:id_post;references:id"`
	Likes     []Like    `gorm:"foreignKey:id_post;references:id"`
	ID        string    `gorm:"column:id;primaryKey;not null"`
	Image     string    `gorm:"column:image;not null"`
	Content   string    `gorm:"column:content;not null"`
	IdUser    string    `gorm:"column:id_user;not null"`
	ImageId   string    `gorm:"column:image_id;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type PostCreate struct {
	Image   string `form:"Image" validate:"max=350"`
	Content string `form:"Content" validate:"required"`
	IdUser  string `form:"IdUser" validate:"required"`
}

type PostUpdate struct {
	Image   string `form:"Image" validate:"max=350"`
	Content string `form:"Content" validate:"required"`
}

type UserPost struct {
	ID       string `json:"id"`
	Image    string `json:"image"`
	Username string `json:"username"`
}
type CommentPost struct {
	ID        int       `json:"id"`
	PostId    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
type LikePost struct {
	IdUser    string    `json:"id_user"`
	CreatedAt time.Time `json:"created_at"`
}

type PostResponse struct {
	ID        string        `json:"id"`
	Image     string        `json:"image"`
	Content   string        `json:"content"`
	User      UserPost      `json:"user"`
	Likes     []LikePost    `json:"like"`
	Comments  []CommentPost `json:"comment"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
