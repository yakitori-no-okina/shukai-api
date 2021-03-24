package user

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type Repository struct {
	database.SqlHandler
}

func (repo *Repository) Store(u domain.UserModel) (id int, err error) {
	result := repo.Create(&u)
	return u.ID, result.Error
}

func (repo *Repository) Update(id int, ue domain.UserForEditting) (err error) {
	user := domain.UserModel{ID: id}
	result := repo.Model(&user).Updates(&ue)
	return result.Error
}

func (repo *Repository) Get(id int) (user *domain.UserModel, err error) {
	var u domain.UserModel
	result := repo.First(&u, id)
	return &u, result.Error
}
