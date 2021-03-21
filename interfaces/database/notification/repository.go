package notification

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type NotificationRepository struct {
	database.SqlHandler
}

func (repo *NotificationRepository) Store(n domain.NotificationModel) (err error) {
	result := repo.Create(&n)
	return result.Error
}

func (repo *NotificationRepository) Getlist() (notifications *domain.Notifications, err error) {
	var notifications domain.Notifications
	result := repo.Find(&notifications)
	return &notifications, result.Error
}
