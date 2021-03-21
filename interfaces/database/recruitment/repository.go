package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type RecruitmentRepository struct {
	database.SqlHandler
}

func (repo *RecruitmentRepository) Store(r domain.RecruitmentModel) (id int, err error) {
	result := repo.Create(&r)
	return &r.ID, result.Error
}

func (repo *RecruitmentRepository) Delete(id int) (err error) {
	result := repo.Delete(&domain.RecruitmentModel{}, id)
	return result.Error
}

func (repo *RecruitmentRepository) Get(id int) (recruitment *domain.RecruitmentModel, err error) {
	var recruitment domain.RecruitmentModel
	result := repo.First(&recruitment, id)
	return &recruitment, result.Error
}
