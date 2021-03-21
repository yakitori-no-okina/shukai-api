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
	return &rc.ID, result.Error
}

func (repo *RecruitmentConditionsRepository) Get(id int) (condition *domain.RecruitmentConditionsModel, err error) {
	var condition domain.RecruitmentConditionsModel
	result := repo.Find(&condition, id)
	return &condition, result.Error
}

func (repo *RecruitmentConditionsRepository) Getlist() (conditionList *[]domain.RecruitmentConditionsModel, err error) {
	var conditionList []domain.RecruitmentConditionsModel
	result := repo.Find(&conditionList)
	return &conditionList, result.Error
}
