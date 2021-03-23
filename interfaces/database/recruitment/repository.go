package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type RecruitmentRepository struct {
	database.SqlHandler
}

func (repo *RecruitmentRepository) Store(r domain.RecruitmentModel) (id uint64, err error) {
	result := repo.Create(&r)
	return r.ID, result.Error
}

func (repo *RecruitmentRepository) Remove(id int) (err error) {
	result := repo.Delete(&domain.RecruitmentModel{}, id)
	return result.Error
}

func (repo *RecruitmentRepository) Get(id int) (recruitment *domain.RecruitmentModel, err error) {
	var r domain.RecruitmentModel
	result := repo.First(&r, id)
	return &r, result.Error
}
