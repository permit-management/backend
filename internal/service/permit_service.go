package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
	"github.com/permit-management/backend/utils"
)

// interface
type PermitService interface {
	CreatePermit(permit *domain.Permit) error
	GetAllPermits() ([]domain.Permit, error)
	GetPermitByID(id uint) (domain.Permit, error)
	UpdatePermit(permit *domain.Permit) error
	DeletePermit(id uint) error
}

// implementation
type permitService struct {
	repo repository.PermitRepository
}

func NewPermitService(repo repository.PermitRepository) PermitService {
	return &permitService{repo: repo}
}

func (s *permitService) CreatePermit(permit *domain.Permit) error {
	// validasi sederhana
	if permit.PermitNumber == "" {
		return errors.New("permit number is required")
	}

	// validasi pekerja (ambil dari relasi Workers)
	if len(permit.Workers) == 0 {
		return errors.New("at least one worker is required")
	}

	worker := permit.Workers[0] // misal ambil worker pertama
	if worker.NIK == "" {
		return errors.New("NIK is required")
	}
	if worker.Email == "" {
		return errors.New("Email is required")
	}

	// set timestamp
	permit.CreatedAt = time.Now()
	permit.UpdatedAt = time.Now()

	// simpan ke DB (sekalian akan simpan workers karena ada relasi)
	if err := s.repo.Create(permit); err != nil {
		return err
	}

	// kirim email notifikasi ke worker
	subject := "Permit Created: " + permit.PermitNumber
	body := "Halo,\r\n\r\n" +
		"Permit Anda dengan nomor " + permit.PermitNumber + " berhasil dibuat.\r\n" +
		"NIK: " + worker.NIK + "\r\n\r\n" +
		"Terima kasih."

	if err := utils.SendEmail(worker.Email, subject, body); err != nil {
		fmt.Println("Gagal kirim email:", err)
		return err
	} else {
		fmt.Println("Email berhasil dikirim ke:", worker.Email)
	}

	return nil
}

func (s *permitService) GetAllPermits() ([]domain.Permit, error) {
	return s.repo.FindAll()
}

func (s *permitService) GetPermitByID(id uint) (domain.Permit, error) {
	if id == 0 {
		return domain.Permit{}, errors.New("invalid ID")
	}
	return s.repo.FindByID(id)
}

func (s *permitService) UpdatePermit(permit *domain.Permit) error {
	if permit.ID == 0 {
		return errors.New("ID is required for update")
	}
	permit.UpdatedAt = time.Now()
	return s.repo.Update(permit)
}

func (s *permitService) DeletePermit(id uint) error {
	if id == 0 {
		return errors.New("invalid ID")
	}
	return s.repo.Delete(id)
}
