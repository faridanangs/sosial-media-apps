package helper

import (
	"github.com/faridanang/jasangku-kodu/model"
)

func NotificationResponse(req model.Notification) model.NotificationResponse {
	return model.NotificationResponse{
		ID:        req.ID,
		PostId:    req.IdPost,
		TypeNotif: req.TypeNotif,
		IsRead:    req.IsRead,
		User: model.UserComment{
			ID:       req.User.ID,
			Image:    req.User.Image,
			Username: req.User.UserName,
		},
		CreatedAt: req.CreatedAt,
	}
}

func NotificationResponses(req []model.Notification) []model.NotificationResponse {
	var notifications []model.NotificationResponse
	for _, u := range req {
		notifications = append(notifications, NotificationResponse(u))
	}
	return notifications
}
