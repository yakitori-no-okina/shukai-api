package action

import (
	"net/http"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/approvalwait"
	"shukai-api/interfaces/database/notification"
	"shukai-api/interfaces/database/recruitment"
	"shukai-api/usecase"
	"strconv"
)

type ApprovalDecideAction struct {
	Interactor usecase.ApprovalInteractor
}

func NewApprovalDecideAction(sqlHandler database.SqlHandler) *ApprovalDecideAction {
	return &ApprovalDecideAction{
		Interactor: usecase.ApprovalInteractor{
			Aw: &approvalwait.Repository{
				SqlHandler: sqlHandler,
			},
			Ru: &recruitment.UsersRepository{
				SqlHandler: sqlHandler,
			},
			R: &recruitment.Repository{
				SqlHandler: sqlHandler,
			},
			N: &notification.Repository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *ApprovalDecideAction) Decide(c Context) error {
	approvalwait_id, _ := strconv.Atoi(c.Param("approvalwait_id"))
	should_approval, _ := strconv.ParseBool(c.Param("should_approval"))

	error_for_decide := action.Interactor.Decide(approvalwait_id, should_approval)
	if error_for_decide != nil {
		return c.JSON(http.StatusBadRequest, error_for_decide)
	} else {
		return c.JSON(http.StatusNoContent, nil)
	}
}
