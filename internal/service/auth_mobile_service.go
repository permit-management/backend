package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
)

type AuthMobileService interface {
	Login(permitNumber string, nik string) (string, *domain.Worker, error)
}

type authMobileService struct {
	authMobileRepo repository.AuthMobileRepository
	jwtSecret      string
}

func NewAuthMobileService(repo repository.AuthMobileRepository, jwtSecret string) AuthMobileService {
	return &authMobileService{authMobileRepo: repo, jwtSecret: jwtSecret}
}

func (s *authMobileService) Login(permitNumber string, nik string) (string, *domain.Worker, error) {
	worker, err := s.authMobileRepo.FindByPermitAndNIK(permitNumber, nik)
	if err != nil {
		return "", nil, errors.New("invalid permit number or NIK")
	}

	// generate JWT
	claims := jwt.MapClaims{
		"worker_id": worker.ID,
		"nik":       worker.NIK,
		"permit_id": worker.PermitID,
		"exp":       time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, worker, nil
}
