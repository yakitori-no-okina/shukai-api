package recruitment

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type ConditionsRepository struct {
	database.SqlHandler
}

func (repo *ConditionsRepository) Store(rc domain.RecruitmentConditionsModel) (id int, err error) {
	result := repo.Create(&rc)
	return rc.ID, result.Error
}

func (repo *ConditionsRepository) Get(id int) (condition *domain.RecruitmentConditionsModel, err error) {
	var c domain.RecruitmentConditionsModel
	result := repo.Find(&c, id)
	return &c, result.Error
}

func (repo *ConditionsRepository) Getlist() (conditionList *[]domain.RecruitmentConditionsModel, err error) {
	var cl []domain.RecruitmentConditionsModel
	result := repo.Find(&cl)
	return &cl, result.Error
}
