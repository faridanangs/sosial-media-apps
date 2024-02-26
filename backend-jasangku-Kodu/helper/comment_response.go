package helper

import (
	"github.com/faridanang/jasangku-kodu/model"
)

func CommentResponse(req model.Comment) model.CommentResponse {
	return model.CommentResponse{
		ID:      req.ID,
		Comment: req.Comment,
		PostId:  req.IdPost,
		User: model.UserComment{
			ID:       req.User.ID,
			Image:    req.User.Image,
			Username: req.User.UserName,
		},
		CreatedAt: req.CreatedAt,
	}
}

func CommentResponses(req []model.Comment) []model.CommentResponse {
	var comments []model.CommentResponse
	for _, u := range req {
		comments = append(comments, CommentResponse(u))
	}
	return comments
}
