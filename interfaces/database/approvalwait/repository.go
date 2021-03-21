package approvalwait

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type ApprovalWaitRepository struct {
	database.SqlHandler
}

func (repo *ApprovalWaitRepository) Store(a domain.ApprovalWaitModel) (err error) {
	result := repo.Create(&a)
	return result.Error
}

func (repo *ApprovalWaitRepository) Delete(recruitment_id, user_id) (err error) {
	result := repo.Where("recruitment_id = ? AND user_id = ?", recruitment_id, user_id).Delete(&domain.ApprovalWaitModel{})
	return result.Error
}
