package user

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type UserRepository struct {
	database.SqlHandler
}

func (repo *UserRepository) Store(ua domain.UserForAdding) (id int, err error) {
	result := repo.Create(&ua)
	return &u.ID, result.Error
}

func (repo *UserRepository) Update(id int, ue domain.UserForEditting) (err error) {
	user := domain.UserModel{ID: id}
	result := repo.Model(&user).Update(&ue)
	return result.Error
}

func (repo *UserRepository) Get(id int) (user *domain.UserModel, err error) {
	var user domain.UserModel
	result := repo.First(&user, id)
	return &user, result.Error
}
