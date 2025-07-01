package integration

import (
	"net/http"
	"time"

	"github.com/permit-management/backend/pkg/logger"
)

type NotificationService interface {
	SendNotification(req SendNotificationRequest) (*NotificationResponse, error)
}

type notificationService struct {
	// config setting.NotificationS
	client http.Client
}

type SendNotificationRequest struct {
	Type        string `json:"type"`
	Destination string `json:"destionation"`
	Title       string `json:"title"`
	Body        string `json:"body"`
}

type NotificationResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"message"`
	Data any       `json:"data,omitempty"`
	Time time.Time `json:"timestamp,omitempty"`
}

func NewNotificationService( /* config setting.NotificationS */ ) NotificationService {
	return &notificationService{ /* config: config, */ client: http.Client{
		Timeout: 5 * time.Second,
	}}
}

func (k *notificationService) SendNotification(req SendNotificationRequest) (*NotificationResponse, error) {
	var err error
	defer func() {
		if err != nil {
			logger.Log().Infof("NotificationService SendNotification Failed %v", err)
		}
	}()

	return nil, nil
}
