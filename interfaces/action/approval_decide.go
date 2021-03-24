package action

import (
	"net/http"
	"shukai-api/interfaces/database/approvalwait"
	"shukai-api/interfaces/database/notification"
	"shukai-api/interfaces/database/recruitment"
	"shukai-api/usecase"
	"strconv"
)

type ApprovalDecideAction struct {
	Interactor usecase.ApprovalInteractor
}

func NewApprovalDecideAction() *ApprovalDecideAction {
	return &ApprovalDecideAction{
		Interactor: usecase.ApprovalInteractor{
			Aw: &approvalwait.Repository{},
			Ru: &recruitment.UsersRepository{},
			N:  &notification.Repository{},
		},
	}
}

func (action *ApprovalDecideAction) Decide(c Context) error {
	approvalwait_id, _ := strconv.Atoi(c.Param("approvalwait_id"))
	should_approval, _ := strconv.ParseBool(c.Param("should_approval"))

	success, error_for_decide := action.Interactor.Decide(approvalwait_id, should_approval)
	if success {
		return nil
	} else {
		return c.JSON(http.StatusBadRequest, error_for_decide)
	}
}
