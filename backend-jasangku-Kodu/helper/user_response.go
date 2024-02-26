package helper

import (
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/gofiber/fiber/v2"
)

func UserResponse(user model.User) model.UserResponse {
	return model.UserResponse{
		ID:        user.ID,
		Image:     user.Image,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserResponses(user []model.User) []model.UserResponse {
	var users []model.UserResponse
	for _, u := range user {
		users = append(users, UserResponse(u))
	}
	return users
}

func WebResponse(c *fiber.Ctx, res any) error {
	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   res,
	})
}
