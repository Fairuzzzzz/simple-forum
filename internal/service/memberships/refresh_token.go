package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/Fairuzzzzz/simpleform/internal/model/memberships"
	"github.com/Fairuzzzzz/simpleform/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	exsistingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err
	}

	// Jika existing token nya sudah expire semua maka kembalikan error
	if exsistingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	// Jika token di database tidak sama dengan request token, kembalikan error invalid refresh
	if exsistingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exists")
	}

	// Jika token di database sama dengan request token maka akan generate jwt
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}
