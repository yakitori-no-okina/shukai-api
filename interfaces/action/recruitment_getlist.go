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

type RecruitmentGetListAction struct {
	Interactor usecase.RecruitmentInteractor
}

func NewRecruitmentGetListAction(sqlHandler database.SqlHandler) *RecruitmentGetListAction {
	return &RecruitmentGetListAction{
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

func (action *RecruitmentGetListAction) GetList(c Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	recruitments, error_for_get := action.Interactor.GetList(user_id)
	if error_for_get != nil {
		return c.JSON(http.StatusBadRequest, error_for_get.Error)
	}

	return c.JSON(http.StatusOK, recruitments)
}
