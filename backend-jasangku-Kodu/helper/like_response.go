package helper

import (
	"github.com/faridanang/jasangku-kodu/model"
)

func LikeResponse(req model.Like) model.LikeResponse {
	return model.LikeResponse{
		IdUser:    req.IdUser,
		IdPost:    req.IdPost,
		CreatedAt: req.CreatedAt,
	}
}

func LikeResponses(req []model.Like) []model.LikeResponse {
	var likes []model.LikeResponse
	for _, u := range req {
		likes = append(likes, LikeResponse(u))
	}
	return likes
}
