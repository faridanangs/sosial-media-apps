package handler

import (
	"github.com/faridanang/jasangku-kodu/helper"
	"github.com/faridanang/jasangku-kodu/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type NotificationHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type NotificationHandlerIPLM struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewNotificationHandler(db *gorm.DB, validate *validator.Validate) NotificationHandler {
	return &NotificationHandlerIPLM{
		DB:       db,
		Validate: validate,
	}
}

func (handler *NotificationHandlerIPLM) Create(ctx *fiber.Ctx) error {
	notificationRequest := model.NotificationCreate{}
	if err := ctx.BodyParser(&notificationRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(notificationRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}

	notification := model.Notification{
		TypeNotif: notificationRequest.TypeNotif,
		IdUser:    notificationRequest.IdUser,
		IdPost:    notificationRequest.IdPost,
		IdFriend:  notificationRequest.IdFriend,
	}
	if err := handler.DB.Debug().Create(&notification).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	notificationResponse := model.Notification{}
	if err := handler.DB.Debug().Preload("User").Model(model.Notification{}).Where("id_friend = ?", notification.IdFriend).Take(&notificationResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}

	return helper.WebResponse(ctx, helper.NotificationResponse(notificationResponse))

}

func (handler *NotificationHandlerIPLM) Update(ctx *fiber.Ctx) error {
	handler.DB.Exec(`
    DELETE FROM notifications
    WHERE id NOT IN (
        SELECT MIN(id) 
        FROM notifications 
        GROUP BY id_friend
    )
`)
	notificationRequest := model.NotificationUpdate{}
	if err := ctx.BodyParser(&notificationRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	if err := handler.Validate.Struct(notificationRequest); err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Bad request", fiber.StatusBadRequest)
	}
	notification := model.Notification{}
	if err := handler.DB.Debug().Model(model.Notification{}).Where("id_friend = ?", notificationRequest.IdFriend).Take(&notification).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Not Found", fiber.StatusNotFound)
	}
	notification.IsRead = notificationRequest.IsRead

	if err := handler.DB.Debug().Where("id_post = ? and id_user = ?", notification.IdPost, notification.IdUser).Updates(&notification).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}
	notificationResponse := model.Notification{}
	if err := handler.DB.Debug().Model(model.Notification{}).Preload("User").Where("id_friend = ?", notification.IdUser).Take(&notificationResponse).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), "Internal Server Error", fiber.StatusInternalServerError)
	}
	return helper.WebResponse(ctx, helper.NotificationResponse(notificationResponse))
}
func (handler *NotificationHandlerIPLM) Delete(ctx *fiber.Ctx) error {
	if err := handler.DB.Debug().Model(model.Notification{}).Where("id_post = ? and id_user = ?", ctx.Params("id"), ctx.Query("q")).Delete(&model.Notification{}).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, "Delete Success")
}

func (handler *NotificationHandlerIPLM) GetByID(ctx *fiber.Ctx) error {
	notifications := []model.Notification{}
	if err := handler.DB.Debug().Preload("User").Model(model.Notification{}).Where("id_user = ?", ctx.Params("id")).Find(&notifications).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.NotificationResponses(notifications))
}
func (handler *NotificationHandlerIPLM) GetAll(ctx *fiber.Ctx) error {
	notifications := []model.Notification{}
	if err := handler.DB.Debug().Preload("User").Model(model.Notification{}).Find(&notifications).Error; err != nil {
		return helper.ErrorResponse(ctx, err.Error(), " NotFound", fiber.StatusNotFound)
	}
	return helper.WebResponse(ctx, helper.NotificationResponses(notifications))
}
