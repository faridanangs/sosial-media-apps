package model

import "time"

type User struct {
	ID        string    `gorm:"column:id;primaryKey;not null"`
	Image     string    `gorm:"column:image"`
	FirstName string    `gorm:"column:first_name;not null"`
	LastName  string    `gorm:"column:last_name;not null"`
	UserName  string    `gorm:"column:username;not null"`
	Email     string    `gorm:"column:email;not null"`
	Password  string    `gorm:"column:password;not null"`
	IsAdmin   bool      `gorm:"column:is_admin;default:false"`
	ImageId   string    `gorm:"column:image_id;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type UserCreate struct {
	ID        string `form:"id"`
	Image     string `form:"Image" validate:"max=350"`
	FirstName string `form:"FirstName" validate:"required,max=20,min=3"`
	LastName  string `form:"LastName" validate:"max=20"`
	UserName  string `form:"UserName" validate:"required,max=30,min=3"`
	Email     string `form:"Email" validate:"required,email,max=120,min=5"`
	Password  string `form:"Password" validate:"required"`
	IsAdmin   bool   `form:"IsAdmin"`
}

type UserUpdate struct {
	Image     string `form:"Image" validate:"max=350"`
	FirstName string `form:"FirstName" validate:"required,max=20,min=3"`
	LastName  string `form:"LastName" validate:"max=20"`
	UserName  string `form:"UserName" validate:"required,max=30,min=3"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Image     string    `json:"image"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
