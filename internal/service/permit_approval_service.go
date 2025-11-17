package service

import (
	"fmt"
	"time"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
	"github.com/permit-management/backend/utils"
)

type PermitApprovalService interface {
	ApprovePermit(approval *domain.PermitApproval) error
}

type permitApprovalService struct {
	repo       repository.PermitApprovalRepository
	permitRepo repository.PermitRepository
}

func NewPermitApprovalService(repo repository.PermitApprovalRepository, permitRepo repository.PermitRepository) PermitApprovalService {
	return &permitApprovalService{repo: repo, permitRepo: permitRepo}
}

func (s *permitApprovalService) ApprovePermit(approval *domain.PermitApproval) error {
	approval.CreatedAt = time.Now()
	approval.UpdatedAt = time.Now()

	// Simpan approval ke DB
	if err := s.repo.Create(approval); err != nil {
		return err
	}

	// Update status permit
	if err := s.permitRepo.UpdateStatus(approval.PermitID, approval.Status); err != nil {
		return err
	}

	// Ambil data permit & workers
	permit, err := s.permitRepo.FindByID(approval.PermitID)
	if err != nil {
		return err
	}

	// Kirim email ke semua worker terkait permit
	for _, worker := range permit.Workers {

		var subject, body string

		// Jika Approved
		if approval.Status == "Approved" {
			subject = "Permit Approved: " + permit.PermitNumber
			body = fmt.Sprintf(
				"Halo %s,\r\n\r\nPermit Anda dengan nomor %s telah DISETUJUI.\r\n"+
					"NIK/KTP: %s\r\n\r\nTerima kasih.",
				worker.Name,
				permit.PermitNumber,
				worker.NIK,
			)
		}

		// Jika Rejected
		if approval.Status == "Rejected" {
			subject = "Permit Rejected: " + permit.PermitNumber
			body = fmt.Sprintf(
				"Halo %s,\r\n\r\nPermit Anda dengan nomor %s telah DITOLAK.\r\n"+
					"Alasan Penolakan: %s\r\n\r\n"+
					"Silakan diperbaiki dan ajukan ulang.\r\n\r\nTerima kasih.",
				worker.Name,
				permit.PermitNumber,
				approval.Note,
			)
		}

		// Execute kirim email
		if subject != "" {
			if err := utils.SendEmail(worker.Email, subject, body); err != nil {
				fmt.Println("Gagal kirim email:", err)
			} else {
				fmt.Println("Email berhasil dikirim ke:", worker.Email)
			}
		}
	}

	return nil
}