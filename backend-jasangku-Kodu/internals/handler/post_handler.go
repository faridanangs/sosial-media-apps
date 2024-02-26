package handler

import (
	"github.com/faridanang/jasangku-kodu/helper"
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByIDPost(ctx *fiber.Ctx) error
}

type PostHandlerIPLM struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewPostHandler(db *gorm.DB, validate *validator.Validate) PostHandler {
	return &PostHandlerIPLM{
		DB:       db,
		Validate: validate,
	}
}

func (handler *PostHandlerIPLM) Create(ctx *fiber.Ctx) error {
	postRequest := model.PostCreate{}
	if err := ctx.BodyParser(&postRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(postRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	fileHeader, err := ctx.FormFile("Image")
	if err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	fileImage, err := fileHeader.Open()
	if err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	cloudynaryUploader := helper.CreateImageToCloudinary(fileImage, ctx)
	post := model.Post{
		ID:      uuid.NewString(),
		Image:   cloudynaryUploader.SecureURL,
		Content: postRequest.Content,
		IdUser:  postRequest.IdUser,
		ImageId: cloudynaryUploader.PublicID,
	}
	if err := handler.DB.Debug().Create(&post).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	PostResponse := model.Post{}
	if err := handler.DB.Debug().Model(model.Post{}).Where("id = ?", post.ID).Take(&PostResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.PostResponse(PostResponse))

}

func (handler *PostHandlerIPLM) Update(ctx *fiber.Ctx) error {
	postRequest := model.PostUpdate{}
	if err := ctx.BodyParser(&postRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(postRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	post := model.Post{}
	if err := handler.DB.Debug().Model(model.Post{}).Where("id = ?", ctx.Params("id")).Take(&post).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}

	helper.DeleteImageInCloudinary(post.ImageId, ctx)

	fileHeader, err := ctx.FormFile("Image")
	if err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	fileImage, err := fileHeader.Open()
	if err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	cloudynaryUploader := helper.CreateImageToCloudinary(fileImage, ctx)

	post.Content = postRequest.Content
	post.Image = cloudynaryUploader.SecureURL
	post.ImageId = cloudynaryUploader.PublicID

	if err := handler.DB.Debug().Updates(&post).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	PostResponse := model.Post{}
	if err := handler.DB.Debug().Model(model.Post{}).Where("id = ?", post.ID).Take(&PostResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}
	return helper.WebResponse(ctx, helper.PostResponse(PostResponse))

}
func (handler *PostHandlerIPLM) Delete(ctx *fiber.Ctx) error {
	if err := handler.DB.Debug().Model(model.Post{}).Where("id = ?", ctx.Params("id")).Delete(&model.Post{}).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, "Delete Success")
}

func (handler *PostHandlerIPLM) GetByID(ctx *fiber.Ctx) error {
	posts := []model.Post{}
	if err := handler.DB.Debug().Preload("User").Preload("Comments").Preload("Likes").Model(model.Post{}).Where("id_user = ?", ctx.Params("id")).Find(&posts).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.PostResponses(posts))
}
func (handler *PostHandlerIPLM) GetByIDPost(ctx *fiber.Ctx) error {
	post := model.Post{}
	if err := handler.DB.Debug().Preload("User").Preload("Comments").Preload("Likes").Model(model.Post{}).Where("id = ?", ctx.Params("id")).Take(&post).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.PostResponse(post))
}
func (handler *PostHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	posts := []model.Post{}
	if err := handler.DB.Debug().Preload("User").Preload("Comments").Preload("Likes").Model(model.Post{}).Find(&posts).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.PostResponses(posts))
}
