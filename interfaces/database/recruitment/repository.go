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
	var rc_models []domain.RecruitmentConditionsModel
	result := repo.Where(
		"backend <= ? AND frontend <= ? AND management <= ? AND mobile <= ? AND ai <= ?",
		skill.Backend,
		skill.Frontend,
		skill.Management,
		skill.Mobile,
		skill.AI,
	).Find(&rc_models)
	if result.Error != nil {
		return []domain.RecruitmentModel{}, result.Error
	}

	var r_ids []int
	for _, rc_model := range rc_models {
		r_ids = append(r_ids, rc_model.RecruitmentID)
	}

	var rs []domain.RecruitmentModel
	result_for_get := repo.Where(r_ids).Find(&rs)
	return rs, result_for_get.Error
}
