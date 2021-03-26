package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type Repository struct {
	database.SqlHandler
}

func (repo *Repository) Store(r *domain.RecruitmentModel) (id int, err error) {
	result := repo.Create(r)
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

func (repo *Repository) GetListWithSkill(skill *domain.UserSkill) (recruitments []domain.RecruitmentModel, err error) {
	var rcs []domain.RecruitmentConditionsModel
	result := repo.Where("backend >= ? AND frontend >= ? AND management >= ? AND mobile >= ? AND ai >= ?", skill.Backend, skill.Frontend, skill.Management, skill.Mobile, skill.AI).Find(&rcs)

	var rs []domain.RecruitmentModel
	for _, rc := range rcs {
		rs = append(rs, rc.RecruitmentModel)
	}
	return rs, result.Error
}
