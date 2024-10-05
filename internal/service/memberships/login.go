package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Fairuzzzzz/simpleform/internal/model/memberships"
	"github.com/Fairuzzzzz/simpleform/pkg/jwt"
	tokenUtil "github.com/Fairuzzzzz/simpleform/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}

	// Checking apakah email ada atau tidak
	if user == nil {
		return "", "", errors.New("email not exists")
	}

	// Checking apakah password sudah sesuai dengan email atau belum
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", errors.New("email or password invalid")
	}

	// Integrasi dengan fungsi generate JWT
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	// Integrasi refresh token
	exsistingToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", "", err
	}

	// Checking apakah existing token ada atau tidak
	if exsistingToken != nil {
		return token, exsistingToken.RefreshToken, nil
	}

	// Membuat refresh token jika tidak ada
	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	// JIka sukses (refresh token ada), insert ke database
	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})

	if err != nil {
		log.Error().Err(err).Msg("error inserting refresh token to database")
		return token, refreshToken, err
	}

	return token, refreshToken, nil
}
