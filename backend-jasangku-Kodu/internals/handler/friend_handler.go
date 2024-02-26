package handler

import (
	"github.com/faridanang/jasangku-kodu/helper"
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FriendHandler interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
}

type FriendHandlerIPLM struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewFriendHandler(db *gorm.DB, validate *validator.Validate) FriendHandler {
	return &FriendHandlerIPLM{
		DB:       db,
		Validate: validate,
	}
}

func (handler *FriendHandlerIPLM) Create(ctx *fiber.Ctx) error {
	friendRequest := model.FriendCreate{}
	if err := ctx.BodyParser(&friendRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(friendRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}

	friend := model.Friend{
		IdUser:       friendRequest.IdUser,
		IdUserFriend: friendRequest.IdUserFriend,
	}
	if err := handler.DB.Debug().Create(&friend).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	friendResponse := model.Friend{}
	if err := handler.DB.Debug().Model(model.Friend{}).Where("id_user = ?", friend.IdUser).Take(&friendResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.FrinedResponse(friendResponse))

}
func (handler *FriendHandlerIPLM) Delete(ctx *fiber.Ctx) error {
	friendRequest := model.FriendCreate{}
	if err := ctx.BodyParser(&friendRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(friendRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}

	if err := handler.DB.Debug().Model(model.Friend{}).Where("id_user_friend = ? and id_user =?", friendRequest.IdUserFriend, friendRequest.IdUser).Delete(&model.Friend{}).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, "Delete Success")
}
func (handler *FriendHandlerIPLM) GetByID(ctx *fiber.Ctx) error {
	FriendResponses := []model.Friend{}
	if err := handler.DB.Debug().Model(model.Friend{}).Preload("User").Where("id_user =?", ctx.Params("id")).Find(&FriendResponses).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.FrinedResponses(FriendResponses))
}
func (handler *FriendHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	FriendResponses := []model.Friend{}
	if err := handler.DB.Debug().Model(model.Friend{}).Preload("User").Find(&FriendResponses).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}
	return helper.WebResponse(ctx, helper.FrinedResponses(FriendResponses))
}
