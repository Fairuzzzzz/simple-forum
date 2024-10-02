package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/Fairuzzzzz/simpleform/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error {
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: request.CommentContent,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment to repository")
		return err
	}
	return nil
}
