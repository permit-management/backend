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

	// simpan approval ke DB
	if err := s.repo.Create(approval); err != nil {
		return err
	}

	// update status permit sesuai approval
	if err := s.permitRepo.UpdateStatus(approval.PermitID, approval.Status); err != nil {
		return err
	}

	// kalau status Approved → kirim email ke worker
	if approval.Status == "Approved" {
		permit, err := s.permitRepo.FindByID(approval.PermitID)
		if err != nil {
			return err
		}

		for _, worker := range permit.Workers {
			subject := "Permit Approved: " + permit.PermitNumber

			body := fmt.Sprintf(
				"Halo %s,\r\n\r\nPermit Anda dengan nomor %s telah DISETUJUI.\r\n",
				worker.Name,
				permit.PermitNumber,
			)
			body += fmt.Sprintf("NIK/KTP: %s\r\n\r\n", worker.NIK)

			body += "Daftar aktivitas:\r\n"
			for i, activity := range permit.Activities {
				body += fmt.Sprintf("%d. %s - %s (Status: %s)\r\n",
					i+1,
					activity.Date.Format("2006-01-02 15:04"),
					activity.Description,
					activity.Status,
				)
			}
			body += "\r\nTerima kasih."

			if err := utils.SendEmail(worker.Email, subject, body); err != nil {
				fmt.Println("Gagal kirim email:", err)
			} else {
				fmt.Println("Email berhasil dikirim ke:", worker.Email)
			}
		}
	}

	return nil
}

if approval.Status == "Rejected" {
    permit, err := s.permitRepo.FindByID(approval.PermitID)
    if err != nil {
        return err
    }

    for _, worker := range permit.Workers {
        subject := "Permit Rejected: " + permit.PermitNumber

        body := fmt.Sprintf(
            "Halo %s,\r\n\r\nPermit Anda dengan nomor %s telah DITOLAK.\r\n",
            worker.Name,
            permit.PermitNumber,
        )
        body += fmt.Sprintf("Alasan Penolakan: %s\r\n", approval.Note)
        body += "\r\nSilakan lengkapi data atau ajukan ulang.\r\n\r\nTerima kasih."

        if err := utils.SendEmail(worker.Email, subject, body); err != nil {
            fmt.Println("Gagal kirim email reject:", err)
        } else {
            fmt.Println("Email reject berhasil dikirim ke:", worker.Email)
        }
    }
}