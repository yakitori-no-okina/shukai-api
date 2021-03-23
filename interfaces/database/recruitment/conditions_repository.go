package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type RecruitmentConditionsRepository struct {
	database.SqlHandler
}

func (repo *RecruitmentConditionsRepository) Store(rc domain.RecruitmentConditionsModel) (id int, err error) {
	result := repo.Create(&rc)
	return int(rc.ID), result.Error
}

func (repo *RecruitmentConditionsRepository) Get(id int) (condition *domain.RecruitmentConditionsModel, err error) {
	var c domain.RecruitmentConditionsModel
	result := repo.Find(&c, id)
	return &c, result.Error
}

func (repo *RecruitmentConditionsRepository) Getlist() (conditionList *[]domain.RecruitmentConditionsModel, err error) {
	var cl []domain.RecruitmentConditionsModel
	result := repo.Find(&cl)
	return &cl, result.Error
}
