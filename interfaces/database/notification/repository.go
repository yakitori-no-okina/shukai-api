package notification

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type Repository struct {
	database.SqlHandler
}

func (repo *Repository) Store(n *domain.NotificationModel) (err error) {
	result := repo.Create(n)
	return result.Error
}

func (repo *Repository) Update(id int, has_read bool) (err error) {
	n_model := &domain.NotificationModel{ID: id}
	result := repo.Model(n_model).Update("has_read", has_read)
	return result.Error
}

func (repo *Repository) Getlist(user_id int) (notifications []domain.NotificationModel, err error) {
	var ns []domain.NotificationModel
	result := repo.Where("user_id = ?", user_id).Find(&ns)
	return ns, result.Error
}
