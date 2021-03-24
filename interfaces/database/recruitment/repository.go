package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type Repository struct {
	database.SqlHandler
}

func (repo *Repository) Store(r domain.RecruitmentModel) (id int, err error) {
	result := repo.Create(&r)
	return r.ID, result.Error
}

func (repo *Repository) Remove(id int) (err error) {
	result := repo.Delete(&domain.RecruitmentModel{}, id)
	return result.Error
}

func (repo *Repository) Get(id int) (recruitment *domain.RecruitmentModel, err error) {
	var r domain.RecruitmentModel
	result := repo.First(&r, id)
	return &r, result.Error
}
