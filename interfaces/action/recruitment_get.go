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

type RecruitmentGetAction struct {
	Interactor usecase.RecruitmentInteractor
}

func NewRecruitmentGetAction(sqlHandler database.SqlHandler) *RecruitmentGetAction {
	return &RecruitmentGetAction{
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

func (action *RecruitmentGetAction) Get(c Context) error {
	recruitment_id, _ := strconv.Atoi(c.Param("recruitment_id"))
	recruitment, error_for_get := action.Interactor.GetDetail(recruitment_id)
	if error_for_get != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, recruitment)
}
