package helper

import (
	"github.com/faridanang/jasangku-kodu/model"
)

func PostResponse(req model.Post) model.PostResponse {
	return model.PostResponse{
		ID:      req.ID,
		Image:   req.Image,
		Content: req.Content,
		User: model.UserPost{
			ID:       req.User.ID,
			Image:    req.User.Image,
			Username: req.User.UserName,
		},
		Likes:     LikesResponses(req.Likes),
		Comments:  CommentsResponses(req.Comments),
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
}

func PostResponses(req []model.Post) []model.PostResponse {
	var posts []model.PostResponse
	for _, u := range req {
		posts = append(posts, PostResponse(u))
	}
	return posts
}

func LikesResponses(req []model.Like) []model.LikePost {
	likes := []model.LikePost{}
	for _, r := range req {
		likes = append(likes, model.LikePost{
			IdUser:    r.IdUser,
			CreatedAt: r.CreatedAt,
		})
	}
	return likes
}
func CommentsResponses(req []model.Comment) []model.CommentPost {
	comments := []model.CommentPost{}
	for _, r := range req {
		comments = append(comments, model.CommentPost{
			ID:        r.ID,
			PostId:    r.IdPost,
			CreatedAt: r.CreatedAt,
		})
	}
	return comments
}
