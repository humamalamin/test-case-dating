package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/humamalamin/test-case-dating/api/domains/interfaces"
	"github.com/humamalamin/test-case-dating/api/repositories/models"
	"github.com/rs/zerolog/log"

	authEntity "github.com/humamalamin/test-case-dating/api/domains/entities/auth"
)

type authRepo struct {
	DB *gorm.DB
}

var (
	code      string
	userModel models.User
)

// Login implements interfaces.AuthRepository.
func (repo *authRepo) Login(ctx context.Context, req *authEntity.Auth) (*authEntity.Auth, error) {
	user, err := repo.getUserByEmail(req.Email)
	if err != nil {
		code := "[Repository] Login - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	userResp := &authEntity.Auth{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Gender:     user.Gender,
		Email:      user.Email,
		Password:   "",
		VerifiedAt: user.VerifiedAt.Local().String(),
	}

	return userResp, nil
}

// Register implements interfaces.AuthRepository.
func (repo *authRepo) Register(ctx context.Context, req *authEntity.Auth) error {
	err := repo.checkEmailExists(ctx, req.Email)
	if err != nil {
		code = "[Repository] Register - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	reqModel := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Gender:    req.Gender,
		Password:  req.Password,
	}

	err = repo.DB.Create(&reqModel).Error
	if err != nil {
		code = "[Repository] Register - 2"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}

func (repo *authRepo) checkEmailExists(ctx context.Context, email string) error {
	var count int64
	err := repo.DB.Model(&userModel).Where("email = ?", email).Count(&count).Error
	if err != nil {
		code = "[Repository] CheckEmailExists - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	if count > 0 {
		code = "[Repository] CheckEmailExists - 2"
		log.Error().Err(err).Msg(code)
		err = errors.New("e-mail exists")
		return err
	}

	return nil
}

func (repo *authRepo) getUserByEmail(email string) (*models.User, error) {
	err := repo.DB.Where("email = ?", email).First(&userModel).Error
	if err != nil {
		code := "[Repository] GetUserByEmail - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	return &userModel, nil
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	repo := new(authRepo)
	repo.DB = db

	return repo
}
