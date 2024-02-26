package helper

import (
	"github.com/faridanang/jasangku-kodu/model"
)

func FrinedResponse(req model.Friend) model.FriendResponse {
	return model.FriendResponse{
		IdUser:     req.IdUser,
		UserFriend: model.UserFriend{ID: req.User.ID, Image: req.User.Image, Username: req.User.UserName},
		CreatedAt:  req.CreatedAt,
	}
}
func FrinedResponses(req []model.Friend) []model.FriendResponse {
	var friends []model.FriendResponse
	for _, u := range req {
		friends = append(friends, FrinedResponse(u))
	}
	return friends
}
