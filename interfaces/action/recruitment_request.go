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

type RecruitmentRequestAction struct {
	Interactor usecase.RecruitmentInteractor
}

func NewRecruitmentRequestAction(sqlHandler database.SqlHandler) *RecruitmentRequestAction {
	return &RecruitmentRequestAction{
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
			U: &user.Repository{
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

func (action *RecruitmentRequestAction) Request(c Context) error {
	recruitment_id, _ := strconv.Atoi(c.Param("recruitment_id"))
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	error_for_request := action.Interactor.Request(recruitment_id, user_id)
	if error_for_request != nil {
		return c.JSON(http.StatusBadRequest, error_for_request.Error)
	}
	return c.JSON(http.StatusCreated, nil)
}
