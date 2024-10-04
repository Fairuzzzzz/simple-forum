package posts

import (
	"context"

	"github.com/Fairuzzzzz/simpleform/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}
	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	}
	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comments by id to database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:          postDetail.ID,
			UserID:      postDetail.UserID,
			Username:    postDetail.Username,
			PostTitle:   postDetail.PostTitle,
			PostContent: postDetail.PostContent,
			PostHashtag: postDetail.PostHashtag,
			IsLiked:     postDetail.IsLiked,
		}, LikeCount: likeCount, Comment: comments,
	}, nil
}
