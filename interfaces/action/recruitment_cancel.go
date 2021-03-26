package action

import (
	"net/http"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/approvalwait"
	"shukai-api/interfaces/database/notification"
	"shukai-api/interfaces/database/recruitment"
	"shukai-api/interfaces/database/user"
	"shukai-api/usecase"
	"strconv"
)

type RecruitmentCancelAction struct {
	Interactor usecase.RecruitmentInteractor
}

func NewRecruitmentCancelAction(sqlHandler database.SqlHandler) *RecruitmentCancelAction {
	return &RecruitmentCancelAction{
		Interactor: usecase.RecruitmentInteractor{
			R: &recruitment.Repository{
				SqlHandler: sqlHandler,
			},
			RU: &recruitment.UsersRepository{
				SqlHandler: sqlHandler,
			},
			RC: &recruitment.ConditionsRepository{
				SqlHandler: sqlHandler,
			},
			US: &user.SkillRepository{
				SqlHandler: sqlHandler,
			},
			AW: &approvalwait.Repository{
				SqlHandler: sqlHandler,
			},
			N: &notification.Repository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *RecruitmentCancelAction) Cancel(c Context) error {
	recruitment_id, _ := strconv.Atoi(c.Param("recruitment_id"))
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	error_for_cancel := action.Interactor.Cancel(recruitment_id, user_id)
	if error_for_cancel != nil {
		return c.JSON(http.StatusBadRequest, error_for_cancel.Error)
	}
	return c.JSON(http.StatusNoContent, nil)
}
