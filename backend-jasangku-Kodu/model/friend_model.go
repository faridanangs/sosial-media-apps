package model

import "time"

type Friend struct {
	User         User      `gorm:"foreignKey:id_user_friend;references:id"`
	IdUser       string    `gorm:"column:id_user;not null"`
	IdUserFriend string    `gorm:"column:id_user_friend;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

type FriendCreate struct {
	IdUser       string `json:"IdUser" validate:"required"`
	IdUserFriend string `json:"IdUserFriend" validate:"required"`
}

type UserFriend struct {
	ID       string `json:"id"`
	Image    string `json:"image"`
	Username string `json:"username"`
}

type FriendResponse struct {
	IdUser     string     `json:"id_user"`
	UserFriend UserFriend `json:"user_friend"`
	CreatedAt  time.Time  `json:"created_at"`
}
