package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
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

	worker := permit.Workers[0] // ambil worker pertama
	if worker.NIK == "" {
		return errors.New("NIK is required")
	}
	if worker.Email == "" {
		return errors.New("Email is required")
	}

	// validasi activity (minimal 1 activity masuk)
	if len(permit.Activities) == 0 {
		return errors.New("at least one activity is required")
	}

	for i, activity := range permit.Activities {
		if activity.Description == "" {
			return fmt.Errorf("activity %d description is required", i+1)
		}
		if activity.Status == "" {
			return fmt.Errorf("activity %d status is required", i+1)
		}
		// set default CreatedAt & UpdatedAt (kalau belum dikirim dari JSON)
		if activity.CreatedAt.IsZero() {
			permit.Activities[i].CreatedAt = time.Now()
		}
		permit.Activities[i].UpdatedAt = time.Now()
	}

	// set status awal permit ke "Pending"
	permit.Status = "Pending"

	// set timestamp permit
	permit.CreatedAt = time.Now()
	permit.UpdatedAt = time.Now()

	// simpan ke DB (sekalian akan simpan workers & activities karena ada relasi)
	if err := s.repo.Create(permit); err != nil {
		return err
	}

	// email akan dikirim saat admin melakukan approval (API Approve Permit)

	fmt.Println("Permit berhasil dibuat dengan status Pending, menunggu approval admin.")

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