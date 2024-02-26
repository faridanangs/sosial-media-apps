package handler

import (
	"github.com/faridanang/jasangku-kodu/helper"
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LikeHandler interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
}

type LikeHandlerIPLM struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewLikeHandler(db *gorm.DB, validate *validator.Validate) LikeHandler {
	return &LikeHandlerIPLM{
		DB:       db,
		Validate: validate,
	}
}

func (handler *LikeHandlerIPLM) Create(ctx *fiber.Ctx) error {
	likeRequest := model.LikeCreate{}
	if err := ctx.BodyParser(&likeRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(likeRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}

	like := model.Like{
		IdUser: likeRequest.IdUser,
		IdPost: likeRequest.IdPost,
	}
	if err := handler.DB.Debug().Create(&like).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	likeResponse := model.Like{}
	if err := handler.DB.Debug().Model(model.Like{}).Where("id_post = ?", like.IdPost).Take(&likeResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.LikeResponse(likeResponse))

}
func (handler *LikeHandlerIPLM) Delete(ctx *fiber.Ctx) error {
	if err := handler.DB.Debug().Model(model.Like{}).Where("id_user = ? and id_post = ?", ctx.Params("id"), ctx.Query("q")).Delete(&model.Like{}).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, "Delete Success")
}
func (handler *LikeHandlerIPLM) GetByID(ctx *fiber.Ctx) error {
	likeResponses := []model.Like{}
	if err := handler.DB.Debug().Model(model.Like{}).Where("id_user = ? and id_post = ?", ctx.Params("id"), ctx.Query("q")).Find(&likeResponses).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.LikeResponses(likeResponses))
}
func (handler *LikeHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	likeResponses := []model.Like{}
	if err := handler.DB.Debug().Model(model.Like{}).Find(&likeResponses).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.LikeResponses(likeResponses))
}
