package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/Fairuzzzzz/simpleform/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("email or username already exists")
	}

	// Hash Password
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Model DB
	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass), // Harus memasukan password yang telah di hash
		Username:  req.Username,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	// Implementasi function CreateUser
	return s.membershipRepo.CreateUser(ctx, model)
}
