package approvalwait

import (
	"log"
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type Repository struct {
	database.SqlHandler
}

func (repo *Repository) Store(a *domain.ApprovalWaitModel) (id int, err error) {
	log.Print(a)
	result := repo.Create(a)
	return a.ID, result.Error
}

func (repo *Repository) Remove(id int) (err error) {
	result := repo.Delete(&domain.ApprovalWaitModel{}, id)
	return result.Error
}

func (repo *Repository) RemoveWithRecruitmentID(recruitment_id int) (err error) {
	result := repo.Where("recruitment_id = ?", recruitment_id).Delete(&domain.ApprovalWaitModel{})
	return result.Error
}

func (repo *Repository) GetProperties(id int) (recruitment_id int, user_id int, err error) {
	var aw domain.ApprovalWaitModel
	result := repo.First(&aw, id)
	return aw.RecruitmentID, aw.UserID, result.Error
}

func (repo *Repository) Get(recruitment_id int, user_id int) (approvalwait *domain.ApprovalWaitModel, err error) {
	var aw domain.ApprovalWaitModel
	result := repo.Where(&domain.ApprovalWaitModel{
		RecruitmentID: recruitment_id,
		UserID:        user_id,
	}).First(&aw)
	return &aw, result.Error
}
