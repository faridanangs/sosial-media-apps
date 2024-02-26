package handler

import (
	"github.com/faridanang/jasangku-kodu/helper"
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CommentHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type CommentHandlerIPLM struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewCommentHandler(db *gorm.DB, validate *validator.Validate) CommentHandler {
	return &CommentHandlerIPLM{
		DB:       db,
		Validate: validate,
	}
}

func (handler *CommentHandlerIPLM) Create(ctx *fiber.Ctx) error {
	commentRequest := model.CommentCreate{}
	if err := ctx.BodyParser(&commentRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(commentRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}

	comment := model.Comment{
		Comment: commentRequest.Comment,
		IdUser:  commentRequest.IdUser,
		IdPost:  commentRequest.IdPost,
	}
	if err := handler.DB.Debug().Create(&comment).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	commentResponse := model.Comment{}
	if err := handler.DB.Debug().Preload("User").Model(model.Comment{}).Where("id = ?", comment.ID).Take(&commentResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.CommentResponse(commentResponse))

}

func (handler *CommentHandlerIPLM) Update(ctx *fiber.Ctx) error {
	commentRequest := model.CommentUpdate{}
	if err := ctx.BodyParser(&commentRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(commentRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}

	comment := model.Comment{}
	if err := handler.DB.Debug().Model(model.Comment{}).Where("id = ?", ctx.Params("id")).Take(&comment).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}

	comment.Comment = commentRequest.Comment

	if err := handler.DB.Debug().Updates(&comment).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	commentResponse := model.Comment{}
	if err := handler.DB.Debug().Model(model.Comment{}).Where("id = ?", comment.ID).Take(&commentResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}
	return helper.WebResponse(ctx, helper.CommentResponse(commentResponse))
}
func (handler *CommentHandlerIPLM) Delete(ctx *fiber.Ctx) error {
	if err := handler.DB.Debug().Model(model.Comment{}).Where("id = ?", ctx.Params("id")).Delete(&model.Comment{}).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, "Delete Success")
}

func (handler *CommentHandlerIPLM) GetByID(ctx *fiber.Ctx) error {
	comments := []model.Comment{}
	if err := handler.DB.Debug().Preload("User").Model(model.Comment{}).Where("id_post = ?", ctx.Params("id")).Find(&comments).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.CommentResponses(comments))
}
func (handler *CommentHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	comments := []model.Comment{}
	if err := handler.DB.Debug().Preload("User").Model(model.Comment{}).Find(&comments).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.CommentResponses(comments))
}
