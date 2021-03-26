package action

import (
	"net/http"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/notification"
	"shukai-api/usecase"
	"strconv"
)

type NotificationReadAction struct {
	Interactor usecase.NotificationInteractor
}

func NewNotificationReadAction(sqlHandler database.SqlHandler) *NotificationReadAction {
	return &NotificationReadAction{
		Interactor: usecase.NotificationInteractor{
			N: &notification.Repository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *NotificationReadAction) Read(c Context) error {
	notification_id, _ := strconv.Atoi(c.Param("notification_id"))
	error_for_read := action.Interactor.Read(notification_id)
	if error_for_read != nil {
		return c.JSON(http.StatusBadRequest, error_for_read.Error)
	}
	return c.JSON(http.StatusNoContent, nil)
}
