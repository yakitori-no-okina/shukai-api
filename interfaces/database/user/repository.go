package user

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type UserRepository struct {
	database.SqlHandler
}

func (repo *UserRepository) Store(u domain.UserModel) (id int, err error) {
	result := repo.Create(&u)
	return int(u.ID), result.Error
}

func (repo *UserRepository) Update(id int, ue domain.UserForEditting) (err error) {
	user := domain.UserModel{ID: uint64(id)}
	result := repo.Model(&user).Updates(&ue)
	return result.Error
}

func (repo *UserRepository) Get(id int) (user *domain.UserModel, err error) {
	var u domain.UserModel
	result := repo.First(&u, id)
	return &u, result.Error
}
