package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Fairuzzzzz/simpleform/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivities(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {

	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivities, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activity from database")
		return err
	}

	if userActivities == nil {
		// create user activity
		if !request.IsLiked {
			return errors.New("anda belum pernah like sebelumnya")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		// update user activity
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity to database")
		return err
	}
	return nil
}
