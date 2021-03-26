package action

import (
	"net/http"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/user"
	"shukai-api/usecase"
)

type LoginAction struct {
	Interactor usecase.LoginInteractor
}

func NewLoginAction(sqlHandler database.SqlHandler) *LoginAction {
	return &LoginAction{
		Interactor: usecase.LoginInteractor{
			U: &user.Repository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *LoginAction) Login(c Context) error {
	mail := c.FormValue("mail")
	pw := c.FormValue("password")
	user_id, err := action.Interactor.Authentication(mail, pw)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, err.Error)
	}

	return_value := map[string]interface{}{"id": user_id, "token": "svasdn"}
	return c.JSON(http.StatusOK, return_value)
}
