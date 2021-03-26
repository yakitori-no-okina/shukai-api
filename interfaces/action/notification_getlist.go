package action

import (
	"net/http"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/notification"
	"shukai-api/usecase"
	"strconv"
)

type NotificationGetListAction struct {
	Interactor usecase.NotificationInteractor
}

func NewNotificationGetListAction(sqlHandler database.SqlHandler) *NotificationGetListAction {
	return &NotificationGetListAction{
		Interactor: usecase.NotificationInteractor{
			N: &notification.Repository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *NotificationGetListAction) GetList(c Context) error {
	user_id, _ := strconv.Atoi(c.Param("user_id"))
	n_models, error_for_get := action.Interactor.GetList(user_id)
	if error_for_get != nil {
		return c.JSON(http.StatusBadRequest, error_for_get.Error)
	}
	return c.JSON(http.StatusOK, n_models)
}
