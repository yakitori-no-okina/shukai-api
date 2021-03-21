package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type RecruitmentUsersRepository struct {
	database.SqlHandler
}

func (repo *RecruitmentUsersRepository) Store(ru domain.RecruitmentUsersModel) (id int, err error) {
	result := repo.Create(&ru)
	return &ru.ID, result.Error
}

func (repo *RecruitmentUsersRepository) Get(recruitment_id int) (users *[]domain.RecruitmentUsersModel, err error) {
	var users []domain.RecruitmentUsersModel
	result := repo.Find(&users, recruitment_id)
	return &users, result.Error
}
