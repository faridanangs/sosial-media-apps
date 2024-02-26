package model

import "time"

type Notification struct {
	User      User      `gorm:"foreignKey:id_user;references:id"`
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;not null"`
	IdUser    string    `gorm:"column:id_user;not null"`
	IdFriend  string    `gorm:"column:id_friend;not null"`
	IdPost    string    `gorm:"column:id_post;not null"`
	TypeNotif string    `gorm:"column:type_notif;not null"`
	IsRead    bool      `gorm:"column:is_read;default:false"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

type NotificationCreate struct {
	IdUser    string `json:"IdUser" validate:"required"`
	IdFriend  string `json:"IdFriend" validate:"required"`
	IdPost    string `json:"IdPost" validate:"required"`
	TypeNotif string `json:"TypeNotif" validate:"required"`
}

type NotificationUpdate struct {
	IsRead   bool   `json:"IsRead" validate:"required"`
	IdUser   string `json:"IdUser" validate:"required"`
	IdFriend string `json:"IdFriend" validate:"required"`
	IdPost   string `json:"IdPost" validate:"required"`
}

type NotificationResponse struct {
	ID        int         `json:"id"`
	PostId    string      `json:"post_id"`
	TypeNotif string      `json:"type_notif" validate:"required"`
	IsRead    bool        `json:"is_read"`
	CreatedAt time.Time   `json:"created_at"`
	User      UserComment `json:"user"`
}
