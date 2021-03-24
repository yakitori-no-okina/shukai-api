package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type UsersRepository struct {
	database.SqlHandler
}

func (repo *UsersRepository) Store(ru domain.RecruitmentUsersModel) (id int, err error) {
	result := repo.Create(&ru)
	return ru.ID, result.Error
}

func (repo *UsersRepository) Get(recruitment_id int) (users *[]domain.RecruitmentUsersModel, err error) {
	var us []domain.RecruitmentUsersModel
	result := repo.Find(&us, recruitment_id)
	return &us, result.Error
}
