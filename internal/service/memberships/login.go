package memberships

import (
	"context"
	"errors"

	"github.com/Fairuzzzzz/simpleform/internal/model/memberships"
	"github.com/Fairuzzzzz/simpleform/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	// Checking apakah email ada atau tidak
	if user == nil {
		return "", errors.New("email not exists")
	}

	// Checking apakah password sudah sesuai dengan email atau belum
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("email or password invalid")
	}

	// Integrasi dengan fungsi generate JWT
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", nil
	}
	return token, nil
}
