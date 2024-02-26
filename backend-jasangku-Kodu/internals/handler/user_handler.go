package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/faridanang/jasangku-kodu/helper"
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/faridanang/jasangku-kodu/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetByUserName(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	CreateToken(ctx *fiber.Ctx) error
	GetByEmail(ctx *fiber.Ctx) error
}

type UserHandlerIPLM struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserHandler(db *gorm.DB, validate *validator.Validate) UserHandler {
	return &UserHandlerIPLM{
		DB:       db,
		Validate: validate,
	}
}

func (handler *UserHandlerIPLM) Create(ctx *fiber.Ctx) error {
	userRequest := model.UserCreate{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(userRequest); err != nil {
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
	user := model.User{
		ID:        uuid.NewString(),
		Image:     cloudynaryUploader.SecureURL,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		UserName:  userRequest.UserName,
		Email:     userRequest.Email,
		Password:  util.GeneratePassword(userRequest.Password),
		IsAdmin:   userRequest.IsAdmin,
		ImageId:   cloudynaryUploader.PublicID,
	}
	if err := handler.DB.Debug().Create(&user).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	userResponse := model.User{}
	if err := handler.DB.Debug().Model(model.User{}).Where("id = ?", user.ID).Take(&userResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.UserResponse(userResponse))

}

func (handler *UserHandlerIPLM) Update(ctx *fiber.Ctx) error {
	userRequest := model.UserUpdate{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(userRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	user := model.User{}
	if err := handler.DB.Debug().Model(model.User{}).Where("id = ?", ctx.Params("id")).Take(&user).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}

	helper.DeleteImageInCloudinary(user.ImageId, ctx)
	fmt.Println(userRequest.Image)

	fileHeader, err := ctx.FormFile("Image")
	if err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	fileImage, err := fileHeader.Open()
	if err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	cloudynaryUploader := helper.CreateImageToCloudinary(fileImage, ctx)

	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.UserName = userRequest.UserName
	user.Image = cloudynaryUploader.SecureURL
	user.ImageId = cloudynaryUploader.PublicID

	if err := handler.DB.Debug().Updates(&user).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	userResponse := model.User{}
	if err := handler.DB.Debug().Model(model.User{}).Where("id = ?", user.ID).Take(&userResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.UserResponse(userResponse))

}
func (handler *UserHandlerIPLM) Delete(ctx *fiber.Ctx) error {
	if err := handler.DB.Debug().Model(model.User{}).Where("id = ?", ctx.Params("id")).Delete(&model.User{}).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, "Delete Success")
}

func (handler *UserHandlerIPLM) GetByUserName(ctx *fiber.Ctx) error {
	user := model.User{}
	if err := handler.DB.Debug().Model(model.User{}).Where("username = ?", ctx.Params("username")).Take(&user).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.UserResponse(user))
}
func (handler *UserHandlerIPLM) GetByEmail(ctx *fiber.Ctx) error {
	user := model.User{}
	if err := handler.DB.Debug().Model(model.User{}).Where("email = ?", ctx.Params("email")).Take(&user).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.UserResponse(user))
}
func (handler *UserHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	users := []model.User{}
	if err := handler.DB.Debug().Model(model.User{}).Find(&users).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.UserResponses(users))
}

func (handler *UserHandlerIPLM) CreateToken(ctx *fiber.Ctx) error {
	tokenRequest := model.CreateToken{}
	if err := ctx.BodyParser(&tokenRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(tokenRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	user := model.User{}
	if err := handler.DB.Debug().Model(user).Where("email = ?", tokenRequest.Email).Take(&user).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tokenRequest.Password)); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}

	claim := model.ClaimToken{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add((7 * 24) * time.Hour).Unix(),
		},
	}

	jwtRes := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := jwtRes.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}
	return helper.WebResponse(ctx, model.TokenResponse{
		ID:        user.ID,
		Image:     user.Image,
		Username:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
	})
}
